FROM golang:alpine
MAINTAINER "HashiCorp Terraform Team <terraform@hashicorp.com>"
ADD . /go/src
WORKDIR /go/src
RUN apk add --update git bash
RUN go build 
RUN go install
