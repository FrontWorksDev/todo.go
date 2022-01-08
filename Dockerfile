FROM golang:1.17.5-alpine

ENV GOPATH=/go
ENV ROOT=/go/src/app
RUN mkdir ${ROOT}
WORKDIR ${ROOT}

ADD . ${ROOT}

RUN apk update && apk upgrade
RUN apk add git zsh vim

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]
