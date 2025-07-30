FROM golang:1.24-alpine

RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copy go.mod first for better caching
COPY go.mod ./
RUN go mod download

# Install Air
RUN go install github.com/air-verse/air@latest

# Copy source code
COPY . .

EXPOSE 8080

CMD ["air"]