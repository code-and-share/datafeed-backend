FROM golang:1.12.5-alpine3.9
#FROM arm32v6/golang:1.12.5-alpine3.9
 
RUN mkdir /app 
 
ADD . /app/ 
 
WORKDIR /app 
 
RUN apk add --update \
    git \
  && go get -u github.com/go-sql-driver/mysql \
  && go build -o main . 
 
ENV DB_NAME Paths
ENV DB_USER root
ENV DB_PASS secret
ENV DB_HOST 0.0.0.0
ENV DB_PORT 3306

CMD ["/app/main"]

