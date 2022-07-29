FROM golang:1.18-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY *.go ./
RUN go build -o /gateway-service-app

EXPOSE 8001
CMD ["/gateway-service-app"]
