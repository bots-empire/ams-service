FROM golang:1.18

WORKDIR /app

COPY . .
RUN go build -buildvcs=false -o /bin/ams-service ./cmd/ams

CMD ["/bin/ams-service"]
