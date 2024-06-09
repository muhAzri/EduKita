# Build stage
FROM golang:1.22-alpine AS build

LABEL maintainer="muhammad.azri.f.s@gmail.com"

# Install necessary build tools
RUN apk --no-cache add gcc g++ make git

WORKDIR /go/src/app

# Copy go.mod and go.sum and download dependencies
COPY ./go.mod ./go.sum ./
RUN go mod download && go mod verify

# Copy the rest of the source code
COPY . .

# Initialize and tidy Go modules
RUN go mod tidy

# Build the application with specific flags
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/web-app ./server.go

# Deploy stage
FROM alpine:3.17

# Install certificates
RUN apk --no-cache add ca-certificates

WORKDIR /usr/bin

# Copy the built application and other necessary files
COPY --from=build /go/src/app/bin/web-app ./web-app
COPY --from=build /go/src/app/edu-kita-firebase-admin.json ./edu-kita-firebase-admin.json
COPY --from=build /go/src/app/.env ./.env

# Expose the port the app runs on
EXPOSE 80

# Set the entrypoint for the container
ENTRYPOINT ["./web-app", "--port", "80"]
