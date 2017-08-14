FROM golang
MAINTAINER Regner Blok-Andersen <shadowdf@gmail.com>

ADD . /go/src/github.com/regner/albiondata-backend
WORKDIR /go/src/github.com/regner/albiondata-backend
RUN go get
RUN go install
ENTRYPOINT /go/bin/albiondata-backend

EXPOSE 8080