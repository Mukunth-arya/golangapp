FROM golang:1.13.0-alpine3.10
RUN  mkdir  /curdapp
ADD  . /curdapp
WORKDIR /curdapp
RUN go build -o main .
ENV MONGO_USER=bXVrdW50aAo=
ENV MONGO_PASS=bXVrdW50aAo=
ENV MONGO_DATABASE=mydb1
ENV PORT=9000

EXPOSE 9000


CMD ["/curdapp/main"]
