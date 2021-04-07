FROM golang:1.16

WORKDIR /go/src

COPY . .

RUN go get ./...

RUN GOOS=linux go build

EXPOSE 8000

CMD ["./main"]