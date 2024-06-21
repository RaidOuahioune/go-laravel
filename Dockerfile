# Base stage
FROM golang:alpine3.20 as base

# Install curl and bash
RUN apk add --no-cache curl bash

# Development stage
FROM base as dev

# Install air
RUN curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | bash

# Set the working directory
WORKDIR /opt/app/api

# Copy go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Run the air binary directly
CMD ["/go/bin/air"]
