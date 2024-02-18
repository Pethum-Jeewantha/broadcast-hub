# Use the official Golang image as the base image
FROM golang:1-alpine3.18 AS build
LABEL authors="PethumJeewantha"

# Set the working directory in the container
WORKDIR /app

# Copy the local code to the container
COPY . .

# Build the Go application
RUN go build -o main ./cmd/server/

# Use a smaller base image for the final image
FROM alpine:latest

# Set the working directory in the container
WORKDIR /app

# Copy the executable from the build image to the smaller base image
COPY --from=build /app/main .

# Expose the port that the application will run on
EXPOSE 3200

# Command to run the application
CMD ["./main"]