FROM ubuntu:latest

ENV GO_VERSION=1.21.5
ENV USER root
ENV UID 0
ENV HOME /root
ENV SHELL /bin/bash

SHELL ["/bin/bash", "-c"]

RUN apt-get -yq update
RUN apt-get install -yq wget git gcc vim sudo net-tools

RUN wget -P /tmp "https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz"

RUN tar -C /usr/local -xzf "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"
RUN rm "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

EXPOSE 9999

WORKDIR $GOPATH
