FROM golang:1.12.13-buster as build

WORKDIR /go/src/github.com/jw3/ppc
COPY . /go/src/github.com/jw3/ppc

# todo;; real package management
RUN go get "github.com/alexandrevicenzi/go-sse"  \
           "github.com/eclipse/paho.mqtt.golang" \
           "github.com/xujiajun/gorouter"

RUN make ppc \
 && mkdir /tmp/dist \
 && cp dist/* /tmp/dist

# ------------------------

FROM centos:7
RUN yum -y install openssh openssh-server openssh-clients openssl-libs
COPY --from=build /tmp/dist /usr/local/bin/
EXPOSE 9000
ENTRYPOINT [ "ppc" ]
