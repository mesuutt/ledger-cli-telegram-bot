FROM golang:1.13.5

RUN mkdir -p /go/src/github.com/mesuutt/teledger
COPY . /go/src/github.com/mesuutt/teledger
WORKDIR /go/src/github.com/mesuutt/teledger
RUN go build .



FROM wernight/phantomjs:2.1
USER root
RUN apt-get update
RUN apt-get install -y imagemagick aha ledger
ADD https://jdebp.eu/Repository/debian/dists/stable/main/binary-amd64/ptyget_6_amd64.deb /tmp/
RUN dpkg -i /tmp/ptyget_6_amd64.deb
COPY --from=0 /go/src/github.com/mesuutt/teledger/teledger .

USER phantomjs
