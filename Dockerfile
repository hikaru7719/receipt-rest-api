FROM golang:1.11

WORKDIR /go/src/app
COPY . .

ENV DOCKERIZE_VERSION v0.6.0
RUN apt-get update && apt-get install -y wget \
 && wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
 && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
 && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
