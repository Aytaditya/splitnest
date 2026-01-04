FROM golang:1.24.4

WORKDIR /app

# Copy only dependency files first 
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the code
COPY . .

# Build once
RUN go build -o app cmd/main.go

EXPOSE 8081

# Run compiled binary
CMD ["./app", "--config", "config/local.yaml"]
