FROM golang:1.20-bullseye

WORKDIR /app
COPY . .
RUN go build -o app
CMD ["./app"]
