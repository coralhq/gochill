FROM golang:1.7.0-alpine
MAINTAINER hiraq "hiraq@onebitmedia.com"

RUN apk add -U git bash && rm -rf /var/cache/apk/*

RUN go get -u github.com/jstemmer/go-junit-report
RUN go get github.com/axw/gocov/gocov
RUN go get github.com/AlekSi/gocov-xml
RUN go get github.com/bitly/go-simplejson 
RUN go get github.com/Sirupsen/logrus

COPY . /go/src/gochill
WORKDIR /go/src/gochill

ENTRYPOINT ["./start.sh"]
