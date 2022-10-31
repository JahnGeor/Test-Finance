FROM golang:1.18

RUN go version

ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o avito-finance ./cmd/main/main.go

CMD ["./avito-finance"]