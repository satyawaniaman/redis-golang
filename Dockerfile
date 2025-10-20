FROM golang:1.21-alpine

WORKDIR /app

# Copy source code
COPY . .

# Build the application
RUN go build -o redis-server main.go

# Expose Redis port
EXPOSE 6379

# Run the server
CMD ["./redis-server"]