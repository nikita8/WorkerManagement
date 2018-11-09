FROM amazonlinux:2 AS gobuild

RUN yum install git -y
RUN amazon-linux-extras install golang1.9

ENV GOPATH=/go
ENV WORKSPACE=$GOPATH/src/worker-management

RUN go get github.com/aws/aws-sdk-go/aws
RUN go get github.com/aws/aws-sdk-go/service/dynamodb
RUN go get github.com/satori/go.uuid

# get granitic dev version
RUN git clone -b dev-1.3.0 --single-branch https://github.com/graniticio/granitic $GOPATH/src/github.com/graniticio/granitic
RUN go install github.com/graniticio/granitic
ENV GRANITIC_HOME=$GOPATH/src/github.com/graniticio/granitic

# install the required packages
RUN go get github.com/graniticio/granitic-yaml
RUN go install github.com/graniticio/granitic-yaml/cmd/grnc-yaml-bind
RUN go install github.com/graniticio/granitic-yaml/cmd/grnc-yaml-project
RUN go install github.com/graniticio/granitic/cmd/grnc-ctl

ENV PATH=$PATH:$GOPATH/bin

WORKDIR $WORKSPACE
ADD . $WORKSPACE

RUN grnc-yaml-bind && go build

FROM amazonlinux:2

RUN mkdir -p /var/app

WORKDIR /var/app

COPY --from=gobuild /go/src/worker-management/ .

EXPOSE 3000

RUN ln -s /dev/stdout access.log

CMD ./worker-management
