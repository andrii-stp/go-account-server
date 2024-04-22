FROM golang:1.21.5-alpine3.19

COPY . /go/src/go-account-server

WORKDIR /go/src/go-account-server

RUN go install

WORKDIR /go/bin

EXPOSE 8080

ENTRYPOINT [ "go-account-server" ]