FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod tidy

COPY ../../server/main.go ./

RUN go build -o server

EXPOSE ${PORT}

ENV PORT ${PORT}

CMD ["/app/server"]

