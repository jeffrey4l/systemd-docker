FROM golang as build

WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./... \
    && go install -v ./...

FROM alpine

RUN apk update \
    && apk upgrade \
    && rm -rf /var/cache/apk/*

COPY --from=build /go/bin/systemd-docker /usr/bin/systemd-docker
