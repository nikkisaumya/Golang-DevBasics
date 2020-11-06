#base image for go application                                                                                                                                                                                  
#FROM golang:alpine
FROM golang:1.15-alpine
#create a dir that will hold our application
RUN mkdir /opt/macaddressfetcher
ADD . /opt/
RUN mkdir /etc/conf
RUN chmod 777 /opt/macaddressfetcher
WORKDIR /opt/macaddressfetcher/run
RUN ls /opt/macaddressfetcher/
RUN cp /opt/macaddressfetcher/conf/config.json /etc/conf/
RUN echo $PATH
 
RUN go build -o /tmp/run .
ENTRYPOINT ["/tmp/run"]


