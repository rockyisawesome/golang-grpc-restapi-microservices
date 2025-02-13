# Define the base image
FROM golang:1.23.5-alpine AS builder

# set the working directory
WORKDIR /app

# copy go.mod and go.sum file
COPY ./go.mod ./go.sum ./

# download go dependencies
RUN go mod tidy

# copy source files
COPY ./ .

# build the go fil
RUN go build -o main

# usse a smaller image to run the app
FROM alpine:latest

# set the working directory
WORKDIR /root/

# copy the compiled go binary from the builder image
COPY --from=builder /app/ .

EXPOSE 8090

CMD ["./main"]