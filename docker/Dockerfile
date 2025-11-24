FROM go:1.25.4 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o conveter ./cmd/


FROM alpine:latest
WORKDIR /app 
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/app .

EXPOSE 8080

CMD [ "./app" ]