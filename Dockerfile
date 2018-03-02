FROM golang:1.9.3

# download and build the project
RUN mkdir -p /go/src/github.com/hakunashida/ushirikina
WORKDIR /go/src/github.com/hakunashida/ushirikina

COPY . /go/src/github.com/hakunashida/ushirikina

RUN go get github.com/codegangsta/gin
RUN go-wrapper download
RUN go-wrapper install

# download and install ipfs
RUN curl -O https://dist.ipfs.io/go-ipfs/v0.4.13/go-ipfs_v0.4.13_linux-amd64.tar.gz
RUN tar xvfz go-ipfs_v0.4.13_linux-amd64.tar.gz
RUN cd go-ipfs && bash install.sh
RUN ipfs init

# run the project
ENV PORT 8080
EXPOSE 3000

CMD gin run