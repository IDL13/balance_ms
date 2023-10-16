FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o balance_ms ./cmd/main.go

CMD ["./balance_ms"]