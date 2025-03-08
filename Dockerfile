# Usa uma imagem oficial do Go para build
FROM golang:1.23 as builder

WORKDIR /app

# Copia os arquivos para dentro do container
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Usa uma imagem mais leve para rodar o binário
FROM alpine:latest

WORKDIR /root/

# Copia o binário gerado na etapa anterior
COPY --from=builder /app/app .

# Expõe a porta do servidor
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./app"]
