FROM golang:1.24.4

WORKDIR /app

# Copy everything
COPY . .

# Download deps
RUN go mod tidy

# Expose service port
EXPOSE 8081


CMD ["go", "run", "cmd/main.go"]