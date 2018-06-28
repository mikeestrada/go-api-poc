FROM golang:1.8

WORKDIR /go/src/go-microservice-poc
COPY . .

CMD ["main"]