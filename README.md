
1. git clone https://github.com/nikkisaumya/Golang-DevBasics/edit/develop-fetcher/macaddressfetcher
2. cd Golang-DevBasics -> which contains in Dockerfile
3. build the docker -> docker build -t macaddressfetcher .
4. Add the api key in config file (path: conf/config.json)
5. docker run macaddressfetcher "44:38:39:ff:ef:57" 


```adding go run outout as well```
root# go run run.go -mac="44:38:39:ff:ef:57"  
2020/11/05 17:10:01 Starting CLI Server
Cumulus Networks, Inc
