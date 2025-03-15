# BUILD STAGE

FROM golang:1.21 AS builder

WORKDIR /app
COPY go.mod go.sum ./
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /kims


# RUNTIME STAGE
FROM alpine:latest
WORKDIR /root/
# Copy only the compiled binary from the builder stage
COPY --from=builder /kims ./compiled/
EXPOSE 8080
CMD [ "./compiled/kims" ]