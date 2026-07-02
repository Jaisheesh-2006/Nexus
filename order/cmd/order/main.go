package main

import (
	"log"
	"time"

	"github.com/Jaisheesh-2006/go-graphql-microservice/order"
	"github.com/Jaisheesh-2006/go-graphql-microservice/rabbitmq"
	"github.com/kelseyhightower/envconfig"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/tinrab/retry"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
	AccountURL  string `envconfig:"ACCOUNT_SERVICE_URL"`
	CatalogURL  string `envconfig:"CATALOG_SERVICE_URL"`
	RabbitmqURL string `envconfig:"RABBITMQ_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r order.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = order.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()

	var amqpChan *amqp.Channel
	conn, ch, err := rabbitmq.Connect(cfg.RabbitmqURL)
	if err == nil {
		defer conn.Close()
		defer ch.Close()
		amqpChan = ch
	} else {
		log.Printf("Could not connect to RabbitMQ: %v", err)
	}

	log.Println("Listening on port 8080...")
	s := order.NewService(r, amqpChan)
	log.Fatal(order.ListenGRPC(s, cfg.AccountURL, cfg.CatalogURL, 8080))
}
