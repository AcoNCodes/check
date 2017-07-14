
FROM golang

ADD . /go/src/github.com/AcoNGame/check

RUN go get ./...

RUN go install github.com/AcoNGame/check

ENTRYPOINT /go/bin/check

EXPOSE 8080