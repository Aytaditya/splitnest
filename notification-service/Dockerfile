FROM golang:1.24.4

WORKDIR /app

# Copy only dependency files first 
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the code
COPY . .

# Build once
RUN go build -o app main.go


# Run compiled binary
CMD ["./app"]
