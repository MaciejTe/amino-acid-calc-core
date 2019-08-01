FROM golang:1.12.6-alpine3.10

COPY . /go-project-template

RUN apk update && apk add git && apk add make && apk add gcc
# RUN sh install_swagger.sh # install latest version of swagger
