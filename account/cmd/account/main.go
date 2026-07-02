package main

import (
	"log"
	"time"

	"github.com/Jaisheesh-2006/go-graphql-microservice/account"
	"github.com/Jaisheesh-2006/go-graphql-microservice/rabbitmq"
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
	RabbitmqURL string `envconfig:"RABBITMQ_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r account.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = account.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()

	// Connect to RabbitMQ
	conn, ch, err := rabbitmq.Connect(cfg.RabbitmqURL)
	if err == nil {
		defer conn.Close()
		defer ch.Close()

		// Declare a queue for this service
		q, err := ch.QueueDeclare(
			"",    // name (empty means auto-generated random name)
			false, // durable
			false, // delete when unused
			true,  // exclusive
			false, // no-wait
			nil,   // arguments
		)
		if err == nil {
			// Bind queue to the exchange
			err = ch.QueueBind(
				q.Name,            // queue name
				"",                // routing key
				"orders_exchange", // exchange
				false,
				nil,
			)
			if err == nil {
				msgs, err := ch.Consume(
					q.Name, // queue
					"",     // consumer
					true,   // auto-ack
					false,  // exclusive
					false,  // no-local
					false,  // no-wait
					nil,    // args
				)
				if err == nil {
					go func() {
						for d := range msgs {
							log.Printf("Received OrderCreated Event! Payload: %s", d.Body)
						}
					}()
				}
			}
		}
	} else {
		log.Printf("Could not connect to RabbitMQ: %v", err)
	}

	log.Println("Listening on port 8080...")
	s := account.NewService(r)
	log.Fatal(account.ListenGRPC(s, 8080))
}
