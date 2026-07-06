FROM golang:1.20-alpine AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/Jaisheesh-2006/go-graphql-microservice
COPY go.mod go.sum ./
COPY vendor vendor
COPY rabbitmq rabbitmq
COPY account account
RUN GO111MODULE=on go build -mod vendor -o /go/bin/app ./account/cmd/account

FROM alpine:3.18
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["app"]
