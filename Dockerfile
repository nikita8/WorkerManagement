FROM golang:alpine AS golibsbuild

RUN apk add git

# GOPATH is already set to /go
ENV WORKSPACE=$GOPATH/src/worker-management
ENV GRANITIC_HOME=$GOPATH/src/github.com/graniticio/granitic

RUN go get github.com/graniticio/granitic
RUN go get github.com/aws/aws-sdk-go/aws
RUN go get github.com/aws/aws-sdk-go/service/dynamodb
RUN go get github.com/satori/go.uuid

# install the required packages
RUN go install github.com/graniticio/granitic/cmd/grnc-bind
RUN go install github.com/graniticio/granitic/cmd/grnc-ctl
RUN go install github.com/graniticio/granitic/cmd/grnc-project

WORKDIR $WORKSPACE
ADD . $WORKSPACE

RUN grnc-bind && go build

FROM alpine

RUN apk add --no-cache ca-certificates

EXPOSE 3000

COPY --from=golibsbuild /go/src/worker-management/ .

CMD ./worker-management
