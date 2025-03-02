FROM golang:1.18-alpine

# Adiciona dependências necessárias
RUN apk add --no-cache gcc musl-dev

# Configura ambiente
WORKDIR /app

# Copia arquivos de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o código fonte
COPY . .

# Compila o código
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app ./cmd/main.go

# Usa uma imagem menor para a execução
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=0 /go/bin/app .

EXPOSE 8080
CMD ["./app"]