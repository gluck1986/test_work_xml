FROM golang:1.20

COPY ./ /go/src/app

WORKDIR /go/src/app/


RUN go build -o ./bin/api ./cmd/api/main.go

EXPOSE 8080

CMD [".bin/api"]