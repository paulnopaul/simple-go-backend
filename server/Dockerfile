FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY main.go ./
RUN go mod tidy
RUN go build -o server

EXPOSE ${PORT}

ENV PORT ${PORT}

CMD ["/app/server"]
