# Etapa 1: build com Go completo
FROM golang:1.24.4 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

# Garante binário LINUX, 100% estático
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

# Etapa 2: imagem final super leve
FROM scratch

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 9999

CMD ["./app"]
