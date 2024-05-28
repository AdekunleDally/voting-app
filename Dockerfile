# syntax=docker/dockerfile:1
# Stage 1: Build the Go application
FROM golang:1.18 as builder

# Set the Current Working Directory inside the container
WORKDIR /voting-app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Stage 2: Run the Go application
FROM gcr.io/distroless/base-debian10

WORKDIR /

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /voting-app/main .

# Command to run the executable
CMD ["./main"]
