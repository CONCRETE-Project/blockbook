# initialize from the image

FROM blockbook-build:latest

RUN add-apt-repository "deb http://mirrors.tencentyun.com/ubuntu/ xenial main restricted universe multiverse"
RUN add-apt-repository "deb http://mirrors.tencentyun.com/ubuntu/ xenial-security main restricted universe multiverse"
RUN add-apt-repository "deb http://mirrors.tencentyun.com/ubuntu/ xenial-updates main restricted universe multiverse"

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y devscripts debhelper make dh-systemd dh-exec && \
    apt-get clean

ADD gpg-keys /tmp/gpg-keys
RUN gpg --batch --import /tmp/gpg-keys/*

ADD build-deb.sh /build/build-deb.sh

WORKDIR /build