# Use official Golang image
FROM golang:1.23.4

# Set working directory
WORKDIR /app

# Copy source code
COPY . .

# Install dependencies
RUN go mod tidy

# Build the application
RUN go build -o main .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]
