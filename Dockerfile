FROM golang:1.22

WORKDIR /app/demo
COPY . .

RUN go build ./cmd

EXPOSE 5555
ENTRYPOINT ["./cmd"]