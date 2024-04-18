FROM golang:1.20.1-alpine
WORKDIR /green-api-nyeltay
COPY . .
RUN go build -o green-api .
CMD ["./green-api"]