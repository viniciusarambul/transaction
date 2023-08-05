FROM golang:1.20
WORKDIR /go/src
CMD["tail", "-f", "dev/null"]