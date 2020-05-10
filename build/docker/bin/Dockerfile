# initialize from the image

FROM kindynos/rocksdb:latest

ENV GOLANG_VERSION=go1.13.8.linux-amd64
ENV GOPATH=/go
ENV PATH=$PATH:$GOPATH/bin
ENV CGO_CFLAGS="-I/opt/rocksdb/include"
ENV CGO_LDFLAGS="-L/opt/rocksdb -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4"

# install and configure go
RUN cd /opt && wget https://storage.googleapis.com/golang/$GOLANG_VERSION.tar.gz && \
    tar xf $GOLANG_VERSION.tar.gz
RUN ln -s /opt/go/bin/go /usr/bin/go
RUN mkdir -p $GOPATH
RUN echo -n "GO version: " && go version
RUN echo -n "GOPATH: " && echo $GOPATH

# install build tools
RUN go get github.com/gobuffalo/packr/...

# download pre-loaded depencencies
RUN \
    cleanup() { rm -rf $GOPATH/src/github.com/grupokindynos ; } && \
    trap cleanup EXIT && \
    mkdir -p $GOPATH/src/github.com/grupokindynos && \
    cd $GOPATH/src/github.com/grupokindynos && \
    git clone https://github.com/grupokindynos/coins-explorer.git && \
    cd coins-explorer && \
    env GO111MODULE=on go mod vendor && \
    cp -r vendor /build/vendor

ADD Makefile /build/Makefile

VOLUME /out

WORKDIR /build