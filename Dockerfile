FROM golang:alpine
RUN mkdir /opt/macaddressfetcher
ADD . /opt/
RUN mkdir /etc/conf
WORKDIR /opt/macaddressfetcher/run
RUN ls /opt/macaddressfetcher/
RUN cp /opt/macaddressfetcher/conf/config.json /etc/conf/

RUN go build -o /tmp/run
CMD ["/tmp/run"]


