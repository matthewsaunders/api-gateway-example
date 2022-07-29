FROM golang:1.18-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY *.go ./
RUN go build -o /number-service-app

EXPOSE 8001
CMD ["/number-service-app"]
