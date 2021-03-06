FROM golang:1.15

WORKDIR /app
COPY . .
RUN mkdir logs
RUN go build -o main main.go
CMD ["./main"]