FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

# Replace "my-app" with your actual binary name
RUN go build -o basic ./main.go  

# Expose the port your application listens on
EXPOSE 8081  
# Replace 8080 with the actual port your app listens on

# Run the compiled binary
# Replace "my-app" with your actual binary name
CMD ["./basic"]  