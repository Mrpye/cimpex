############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

#FROM golang:1.17.2-stretch AS builder
# Install git.
# Git is required for fetching the dependencies.

WORKDIR $GOPATH/src/mypackage/
COPY . ./cimpex

WORKDIR $GOPATH/src/mypackage/cimpex/
# Fetch dependencies.
# Using go get.
RUN apk update \                                                                                                                                                                                                                        
  && apk add ca-certificates zip wget tar curl gcc musl-dev linux-headers\                                                                                                                                                                                                      
  && update-ca-certificates

RUN go install
RUN go get -d -v
# Build the binary.
RUN CGO_ENABLED=0  GOARCH=386 GOOS=linux go build -o /go/bin/cimpex
RUN cd /go/bin/
RUN mkdir /go/bin/images
WORKDIR /go/bin/

############################
# STEP 2 build a small image
############################
#FROM ubuntu 
FROM gcr.io/distroless/static
#FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/cimpex /go/bin/cimpex

EXPOSE 8000

VOLUME [ "/go/bin/package_tmp" ]

ENV BASE_FOLDER=/go/bin/images
ENV WEB_IP=localhost
ENV WEB_PORT=8080

ENTRYPOINT ["/go/bin/cimpex","web"]