FROM golang:1.19-alpine

RUN go version
ENV GOPATH=/

COPY ./ ./

EXPOSE ${HTTP_PORT}

RUN go mod download
RUN go build -o weather-observer ./cmd/main.go

CMD ["./weather-observer"]