FROM golang:1.19-alpine

RUN mkdir /srv/go-http-server-template
WORKDIR /srv/go-http-server-template

RUN apk update && \
    apk add git alpine-sdk build-base

RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

RUN go install -a github.com/cosmtrek/air@latest

RUN go install github.com/rubenv/sql-migrate/...@latest

RUN go env -w GOPRIVATE="github.com/myuoncorp"

CMD ["air"]
