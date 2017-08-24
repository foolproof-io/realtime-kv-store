FROM golang:1.8-alpine

ADD . /go/src/hello

RUN go install hello

EXPOSE 80

CMD ["/go/bin/hello"]
