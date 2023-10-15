FROM golang:latest as builder
WORKDIR /app
COPY ./user_service/go.mod ./user_service/go.sum ./
RUN go mod download -x
COPY ./user_service ./
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
WORKDIR /root
COPY --from=builder /app/main .
# This is not necessary
# EXPOSE 8000
CMD ["./main"]