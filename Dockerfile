FROM golang:1.8.3
RUN mkdir -p /go/src/ushirikina
WORKDIR /go/src/ushirikina
COPY . /go/src/ushirikina
RUN go get github.com/codegangsta/gin
RUN go-wrapper download
RUN go-wrapper install
ENV PORT 8000
EXPOSE 3000
CMD gin run