FROM golang:1.16.5-alpine3.13

RUN mkdir /app
ADD go.mod /app/
ADD go.sum /app/
RUN go mod tidy
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
