#FROM golang:1.15.2-alpine
#
#RUN apk update && apk add git
#
#ENV GO111MODULE=on
#
#ENV GOPATH /go
#ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
#
#RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
#WORKDIR $GOPATH/src/backend-github-trending
#
#COPY . .
#
#RUN go mod init backend-github-trending
#
#WORKDIR cmd/pro
#RUN GOOS=linux go build -o app
#
#RUN chmod +x ./app
#ENTRYPOINT ["./app"]
#
#EXPOSE 3000


FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git

ENV GO111MODULE=on

WORKDIR $GOPATH/src/backend-github-trending/
COPY . .

RUN go mod init backend-github-trending

WORKDIR cmd/pro
RUN go build -o /go/bin/hello

FROM alpine:3.4
COPY --from=builder /go/bin/hello /go/bin/hello

RUN chmod +x /go/bin/hello
ENTRYPOINT ["/bin/sh","-c","/go/bin/hello"]