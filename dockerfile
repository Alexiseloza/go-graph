FROM golang:1.21.4-alpine 

WORKDIR $GOPATH/server 
COPY .  .

RUN go get -d -v ./...

RUN go run github.com/99designs/gqlgen generate

RUN go build -o api .

EXPOSE 9000

CMD [ "./api" ]