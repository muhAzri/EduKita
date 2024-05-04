
# build stage
FROM golang:1.22-alpine AS build

LABEL maintainer "muhammad.azri.f.s@gmail.com"

WORKDIR /app

COPY ./go.mod ./go.sum ./

COPY ./.env ./.env

RUN go mod download

COPY ./ ./

RUN go build -o ./app ./server.go

# Run the tests in the container
FROM build AS run-test
RUN go test -v ./...

# deploy stage
FROM golang:1.22-alpine


WORKDIR /app

COPY --from=build ./app ./

EXPOSE 8080

ENTRYPOINT ["./app"]
