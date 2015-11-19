FROM golang:1.4.3-onbuild
MAINTAINER James Norman

WORKDIR /go/src/app

RUN go get github.com/beego/bee

RUN bee pack

CMD bee run