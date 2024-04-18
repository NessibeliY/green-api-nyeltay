FROM golang:1.20.1-alpine
WORKDIR /green-api-nyeltay
COPY . .
RUN go build -o green-api .
EXPOSE 8080
CMD ["./green-api"]
