FROM amazonlinux:2 AS gobuild

RUN yum install git -y
RUN amazon-linux-extras install golang1.9

ENV GOPATH=/go
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

ENV PATH=$PATH:$GOPATH/bin

WORKDIR $WORKSPACE
ADD . $WORKSPACE

RUN grnc-bind && go build

FROM amazonlinux:2

RUN mkdir -p /var/app

WORKDIR /var/app

COPY --from=gobuild /go/src/worker-management/ .

EXPOSE 3000


CMD ./worker-management
