FROM golang:latest

ENV GOPATH /gopath

ENV GIN_MODE release
ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH

ADD sweetcook-backend /gopath/bin/sweetcook-backend
ADD public /gopath/bin/public
ADD config /gopath/bin/config

WORKDIR /gopath/bin
CMD ./sweetcook-backend