FROM golang:1.11

WORKDIR /go/src/app
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
