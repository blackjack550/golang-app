FROM google/golang
MAINTAINER Jack.Wong "huangshuo@intra.nsfocus.com"

# Build app
WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/src/golang-app

RUN go get -t golang-app
RUN go install golang-app

# EXPOSE 80
CMD ["/gopath/bin/golang-app"]
