FROM golang:1.23

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o main ./main.go

CMD ["/app/main"]