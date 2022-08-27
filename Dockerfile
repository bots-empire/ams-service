FROM golang:1.17

WORKDIR /app

COPY . .
RUN go build -o /bin/ams-service ./cmd/ams

CMD ["/bin/ams-service"]
