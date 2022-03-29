# syntax=docker/dockerfile:1

# First, retrieve the most recent go version from the collection of go Docker images
FROM golang:1.17-alpine

# Afterwards, create the working directory into the image
WORKDIR /app

# Copy over the mod and sum files over into the image before downloading the project dependencies.
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

ENV PORT=8080

RUN go build -o bin/app

EXPOSE 8080

CMD [ "/app" ]