FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY .env ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /car-rental-service ./cmd/

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /car-rental-service .

# Создаем директорию для данных
RUN mkdir -p /data
EXPOSE 8080
CMD ["./car-rental-service"]


# docker run --name car-app -p 8080:8080 -d car-rental-service
# docker run -d --name car-app -p 8080:8080 -v car-rental-data:/data car-rental-service
