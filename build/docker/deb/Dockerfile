# initialize from the image

FROM blockbook-build:latest

RUN rm /etc/apt/sources.list && \
    rm /etc/apt/sources.list.* && \
    rm /etc/apt/sources.list.d/* && \
    touch /etc/apt/sources.list && \
    echo "deb http://mirrors.tencentyun.com/ubuntu/ bionic main restricted universe multiverse" >> /etc/apt/sources.list && \
    echo "deb http://mirrors.tencentyun.com/ubuntu/ bionic-security main restricted universe multiverse" >> /etc/apt/sources.list&& \
    echo "deb http://mirrors.tencentyun.com/ubuntu/ bionic-updates main restricted universe multiverse" >> /etc/apt/sources.list&& \
    echo "deb-src http://mirrors.tencentyun.com/ubuntu/ bionic main restricted universe multiverse" >> /etc/apt/sources.list&& \
    echo "deb-src http://mirrors.tencentyun.com/ubuntu/ bionic-security main restricted universe multiverse" >> /etc/apt/sources.list

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y devscripts debhelper make dh-systemd dh-exec && \
    apt-get clean

ADD gpg-keys /tmp/gpg-keys
RUN gpg --batch --import /tmp/gpg-keys/*

ADD build-deb.sh /build/build-deb.sh

WORKDIR /build