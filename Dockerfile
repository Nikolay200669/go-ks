FROM golang:latest
LABEL authors="nik"

# This dockerfile for run test Golang app
# Copy all files from current directory to /app in container
COPY . /app

# Set working directory
WORKDIR /app

# Run go test
CMD ["go", "test", "-v"]

# console run command
# docker build -t go-test .
