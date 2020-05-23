FROM golang:1.14.3-alpine3.11
WORKDIR /app

COPY main.go .
RUN go build -o hello

CMD ["/app/hello"]
