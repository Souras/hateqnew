FROM golang:alpine AS builder

# RUN apk --update add --no-cache ca-certificates openssl git tzdata && update-ca-certificates

# COPY --from=builder /etc/ssl/certs/ca-certificates.crt ./etc/ssl/certs/
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

# Replace "my-app" with your actual binary name
RUN go build -o websocketAPI ./main.go  

# Expose the port your application listens on
EXPOSE 8081  
# Replace 8080 with the actual port your app listens on

# Run the compiled binary
# Replace "my-app" with your actual binary name
CMD ["./websocketAPI"]  