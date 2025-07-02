FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

# Instala as ferramentas necessárias
RUN apk --no-cache add ca-certificates postgresql-client

WORKDIR /root/

# Copia o binário da aplicação
COPY --from=builder /app/main .

# Copia o script de espera
COPY wait-for-db.sh .
RUN chmod +x wait-for-db.sh

EXPOSE 8080 9090 8081

# Usa o script de espera como entrypoint
CMD ["./wait-for-db.sh"] 