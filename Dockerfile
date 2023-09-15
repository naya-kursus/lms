FROM golang:1.21-alpine AS builder

WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .

RUN go mod download -x

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine
WORKDIR /usr/local/bin

COPY --from=builder /go/src/app/views ./views
COPY --from=builder /go/src/app/public ./public
COPY --from=builder /go/src/app/main ./main
RUN chmod +x main

EXPOSE 8080
CMD ["main"]
