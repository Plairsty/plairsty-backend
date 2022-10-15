FROM golang:1.18
WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o /usr/local/bin/server cmd/api/server/main.go

CMD ["/usr/local/bin/server"]