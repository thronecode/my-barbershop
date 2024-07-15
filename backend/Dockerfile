FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN cd cmd/server && go build -o server

EXPOSE 4002

CMD ["./cmd/server/server"]