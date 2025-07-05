FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod ./go.mod

RUN go mod download

COPY . .

RUN go build -o stress-test .

ENTRYPOINT ["./stress-test"]
CMD []
