FROM golang:1.7-onbuild
RUN apt-get update
RUN go get github.com/codegangsta/gin
RUN go get github.com/tools/godep
RUN go get gopkg.in/mgo.v2 & go get gopkg.in/mgo.v2/bson
RUN godep save
ADD ./docker/launch.sh /var/run/launch.sh
RUN chmod g+x /var/run/launch.sh

ENTRYPOINT /var/run/launch.sh