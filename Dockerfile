FROM golang:1.7-onbuild
RUN apt-get update
RUN go get github.com/Masterminds/glide
RUN go get gopkg.in/mgo.v2 & go get gopkg.in/mgo.v2/bson
RUN go get github.com/go-playground/lars & go get net/http
RUN go get time & go get log & go get net/http & go get os
RUN glide init -y
RUN glide i
RUN go get github.com/codegangsta/gin
ADD ./docker/launch.sh /var/run/launch.sh
RUN chmod g+x /var/run/launch.sh

ENTRYPOINT /var/run/launch.sh