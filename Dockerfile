FROM golang:1.17

WORKDIR /app/pair-storing
COPY . .
RUN go build -o ./out/pair-storing ./cmd/server

EXPOSE 9000
CMD ["./out/pair-storing"]

