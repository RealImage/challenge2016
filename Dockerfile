# Use the official Go Docker image as the base
FROM golang:latest


# Set the working directory inside the container
WORKDIR /app

# Copy the source code to the container
COPY . .

# Build the Go application inside the container
RUN go build -o app .

# Set the entry point to run the Go application
CMD ["./app"]