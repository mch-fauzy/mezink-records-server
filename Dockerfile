FROM golang:1.22

RUN go install github.com/rubenv/sql-migrate/...@latest

RUN wget -qO- https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-linux-amd64-v0.6.1.tar.gz | tar xvz -C /usr/local/bin

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

COPY start.sh .
RUN chmod +x start.sh

ENTRYPOINT ["./start.sh"]