# initialize from the image

FROM kindynos/rocksdb:latest

ENV GOLANG_VERSION=go1.14.2.linux-amd64
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

# manually download deps
RUN cd /opt/go/src && \
    mkdir -p golang.org/x && \
    cd golang.org/x && \
    git clone https://github.com/golang/sys.git && \
    git clone https://github.com/golang/crypto.git && \
    git clone https://github.com/golang/sync.git && \
    git clone https://github.com/golang/tools.git

# install build tools
RUN go get github.com/gobuffalo/packr/...

# download pre-loaded depencencies
RUN \
    cleanup() { rm -rf $GOPATH/src/github.com/CONCRETE-Project ; } && \
    trap cleanup EXIT && \
    mkdir -p $GOPATH/src/github.com/CONCRETE-Project && \
    cd $GOPATH/src/github.com/CONCRETE-Project && \
    git clone https://github.com/CONCRETE-Project/blockbook.git && \
    cd blockbook && \
    env GO111MODULE=on go mod vendor && \
    cp -r vendor /build/vendor

ADD Makefile /build/Makefile

VOLUME /out

WORKDIR /build