FROM golang:alpine

RUN apk add git

RUN export GOPATH=$HOME
ENV WORKSPACE=$GOPATH/src/worker-management
ENV GRANITIC_HOME=$GOPATH/src/github.com/graniticio/granitic

WORKDIR $WORKSPACE

ADD . $WORKSPACE

RUN cd $WORKSPACE

RUN ./packages

RUN grnc-bind

RUN go build

ENTRYPOINT ./worker-management
