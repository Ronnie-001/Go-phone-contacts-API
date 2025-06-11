FROM golang:1.24.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /phone-contacts-api 

FROM scratch

COPY --from=builder /phone-contacts-api /phone-contacts-api

EXPOSE 8080

CMD ["/phone-contacts-api"]
