FROM golang:1.21-alpine3.17

RUN addgroup -S salesDash && adduser -S salesDash -G salesDash

WORKDIR /server

COPY . .

RUN go build -o server ./main.go

RUN chown -R salesDash:salesDash /server

EXPOSE 8080

CMD ["./server"]