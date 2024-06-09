# Build stage
FROM golang:1.22-alpine AS build

LABEL maintainer="muhammad.azri.f.s@gmail.com"

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download && go mod verify

COPY . ./


# Build the application
RUN go build -o app ./server.go

# Deploy stage
FROM golang:1.22-alpine

WORKDIR /app

COPY --from=build /app/app ./
COPY --from=build /app/edu-kita-firebase-admin.json ./edu-kita-firebase-admin.json
COPY --from=build /app/.env ./.env

EXPOSE 8080

ENTRYPOINT ["./app"]