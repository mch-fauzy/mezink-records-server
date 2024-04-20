FROM golang:1.22

RUN go install github.com/rubenv/sql-migrate/...@latest

RUN wget -qO- https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-linux-amd64-v0.6.1.tar.gz | tar xvz -C /usr/local/bin

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

RUN chmod +x start.sh

ENTRYPOINT ["./start.sh"]
