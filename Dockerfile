#------
FROM golang:alpine as govendor
WORKDIR /go/src/app
RUN apk add --no-cache git openssl bzr \
    && go get -u github.com/kardianos/govendor
COPY . .
RUN govendor sync
RUN go build

#------
FROM golang:1.13.3-alpine3.10
WORKDIR /go/src/app
COPY --from=govendor /go/src/app/app .
EXPOSE 8080
CMD ["./app"]