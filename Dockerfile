FROM golang:1.20.7-bookworm

WORKDIR /dchants

COPY . .

RUN go build -o dchants

CMD ["./dchants"]

EXPOSE 8080