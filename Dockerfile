# FROM golang:latest AS builder
# WORKDIR /build 

# COPY . .
# RUN go mod download
# RUN go build -o ./conveter

# FROM alpine:latest


# WORKDIR /app
# # SSL certs (IMPORTANT for HTTPS)
# #RUN apk add --no-cache ca-certificatess

# COPY --from=builder /build/conveter ./conveter

# CMD ["/app/conveter"]
FROM golang:latest as build

WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download the Go module dependencies
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o conveter


FROM alpine:latest as run

WORKDIR /app
# Copy the application executable from the build image
COPY --from=build /app/conveter /app/conveter


EXPOSE 8080
CMD ["./conveter"]
