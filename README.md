
1. git clone https://github.com/nikkisaumya/Golang-DevBasics.git
2. cd Golang-DevBasics -> which contains in Dockerfile
3. Add the api key in config file (path: conf/config.json)
4. build the docker -> docker build -t macaddressfetcher .
5. docker run -it macaddressfetcher -mac="44:38:39:ff:ef:57"

**Running on docker**


**#Step1: Docker build macaddressfetcher app**
```
root# docker build -t macaddressfetcher .
Sending build context to Docker daemon  120.8kB
Step 1/11 : FROM golang:1.15-alpine
 ---> d099254f5fc3
Step 2/11 : RUN mkdir /opt/macaddressfetcher
 ---> Using cache
 ---> 0df1f80afc5a
Step 3/11 : ADD . /opt/
 ---> 2ee798f72459
Step 4/11 : RUN mkdir /etc/conf
 ---> Running in cd1d501f2aae
Removing intermediate container cd1d501f2aae
 ---> 0023f8df6dac
Step 5/11 : RUN chmod 777 /opt/macaddressfetcher
 ---> Running in bc421ed9f8ab
Removing intermediate container bc421ed9f8ab
 ---> 48b106a78cec
Step 6/11 : WORKDIR /opt/macaddressfetcher/run
 ---> Running in 4a4d56be712f
Removing intermediate container 4a4d56be712f
 ---> 7c6f47fbdfd2
Step 7/11 : RUN ls /opt/macaddressfetcher/
 ---> Running in f43dd578653c
README
api
cli
conf
go.mod
run
Removing intermediate container f43dd578653c
 ---> 07b2346a5e48
Step 8/11 : RUN cp /opt/macaddressfetcher/conf/config.json /etc/conf/
 ---> Running in 4bea048b49c3
Removing intermediate container 4bea048b49c3
 ---> dad40876d325
Step 9/11 : RUN echo $PATH
 ---> Running in 7b3aac436936
/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
Removing intermediate container 7b3aac436936
 ---> 3bd11b0fb387
Step 10/11 : RUN go build -o /tmp/run .
 ---> Running in 1c12ad825642
Removing intermediate container 1c12ad825642
 ---> 587832db8047
Step 11/11 : ENTRYPOINT ["/tmp/run"]
 ---> Running in b5d2c27d1f01
Removing intermediate container b5d2c27d1f01
 ---> 0dd4bf1d7fea
Successfully built 0dd4bf1d7fea
Successfully tagged macaddressfetcher:latest
```

**#Step2: Docker Run macaddressfetcher app**
```
root# docker run -it macaddressfetcher -mac="44:38:39:ff:ef:57"
```
```
2020/11/06 07:24:28 Starting CLI Server
Cumulus Networks, Inc
```

**adding go run outout as well from terminal**

```
root# go run run.go -mac="44:38:39:ff:ef:57"  
2020/11/05 17:10:01 Starting CLI Server
Cumulus Networks, Inc
```
