FROM golang:1.9.3

RUN mkdir -p /go/src/github.com/hakunashida/ushirikina
WORKDIR /go/src/github.com/hakunashida/ushirikina

COPY . /go/src/github.com/hakunashida/ushirikina

RUN go get github.com/codegangsta/gin
RUN go-wrapper download
RUN go-wrapper install

ENV PORT 8080
EXPOSE 3000

CMD gin run