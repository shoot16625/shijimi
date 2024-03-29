# heroku
FROM golang:1.12.12-alpine3.9 AS build

ENV GOPATH $GOPATH:/go
ENV PATH $PATH:$GOPATH/bin
ENV GO111MODULE=on

RUN apk update && \
  apk add --no-cache git ca-certificates && \
  go get "github.com/go-sql-driver/mysql@v1.5.0" && \
  go get "github.com/beego/bee@v1.11.1" && \
  go get "github.com/astaxie/beego@v1.12.2" && \
  go get "github.com/astaxie/beego/session" && \
  go get "golang.org/x/crypto/bcrypt" && \
  go get "github.com/PuerkitoBio/goquery" && \
  go get "github.com/microcosm-cc/bluemonday" && \
  go get "github.com/ChimeraCoder/anaconda" && \
  rm -rf /var/cache/apk/*

ADD . /go/src/app
WORKDIR /go/src/app

RUN go build -o app main.go

FROM alpine:3.9
RUN apk update && \
apk --no-cache add tzdata ca-certificates && \
  rm -rf /var/cache/apk/*
COPY --from=build /go/src/app/ .
CMD ["./app"]