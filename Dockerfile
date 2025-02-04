# Use official Go image as the base image
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and source code into the container
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your Go project files
COPY . .

# Build the Go application
RUN go build -o app .

# Expose the port from the container
EXPOSE ${PORT}

# Accept build-time arguments
ARG PROXY_URI
ARG PORT=3000

# Set the environment variables
ENV PORT=${PORT}
ENV PROXY_URI=${PROXY_URI}

# Run the Go application
CMD ["./app"]
