FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o myapp

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/myapp .

EXPOSE 8080

CMD ["./myapp"]