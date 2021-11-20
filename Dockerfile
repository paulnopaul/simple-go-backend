FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY go mod download ./

COPY main.go ./

RUN go build -o server

EXPOSE 8081

ENV PORT 8091

CMD ["/app/server"]

