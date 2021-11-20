FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod tidy

COPY main.go ./

RUN go build -o server

EXPOSE 8091

ENV PORT 8091

CMD ["/app/server"]

