FROM golang:1.13.4-buster

RUN mkdir /go/src/prueba

ADD .  /go/src/prueba

WORKDIR  /go/src/prueba

RUN go mod tidy
RUN go get -u gopkg.in/mgo.v2/bson


RUN go build -o /app/server .

EXPOSE 3000


CMD ["/app/server"]