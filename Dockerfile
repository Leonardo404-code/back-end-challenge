FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod /app
COPY go.sum /app
COPY . .

RUN apk add --update --no-cache gcc git musl-dev
RUN go build -ldflags="-s" -o frete-rapido ./cmd/server/main.go
RUN apk add --update --no-cache upx
RUN upx -9 -k frete-rapido

FROM alpine

WORKDIR /app
COPY --from=builder /app/frete-rapido .
COPY --from=builder /app/go.mod /app/go.mod 
COPY --from=builder /app/go.sum /app/go.sum 

EXPOSE 3000

ENTRYPOINT [ "/app/frete-rapido" ]
