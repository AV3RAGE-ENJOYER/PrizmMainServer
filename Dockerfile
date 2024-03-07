# syntax=docker/dockerfile:1

FROM golang:1.21.4
WORKDIR /server
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN export GOOS=linux
RUN go build -o /server
EXPOSE 8080
CMD ["/server"]