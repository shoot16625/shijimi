FROM golang:1.12.12-alpine3.9 AS build

ENV GOPATH $GOPATH:/go
ENV PATH $PATH:$GOPATH/bin

RUN apk update && \
  apk add --no-cache git && \
  go get "github.com/go-sql-driver/mysql" && \
  go get "github.com/beego/bee" && \
  go get "github.com/astaxie/beego" && \
  go get "github.com/astaxie/beego/session" && \
  go get "golang.org/x/crypto/bcrypt" && \
  go get "github.com/PuerkitoBio/goquery" && \
  go get "github.com/microcosm-cc/bluemonday"

ADD . /go/src/app
WORKDIR /go/src/app

#RUN go build -o app main.go
#CMD ["./app"]