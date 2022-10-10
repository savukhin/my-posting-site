FROM golang:1.19.1-alpine

WORKDIR $GOPATH/my-posting-site/backend
COPY . .

RUN sh ./common/protobuf/generate.sh

EXPOSE 3000 
EXPOSE 3100 
EXPOSE 3200
EXPOSE 3400 

ENTRYPOINT ["sh", "./run.sh"]