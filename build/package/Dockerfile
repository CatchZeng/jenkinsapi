FROM golang:1.17.6-alpine as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go mod tidy && go build -o jenkinsapi main.go


FROM alpine:3.7

COPY --from=builder /build/jenkinsapi /usr/local/bin/jenkinsapi
RUN chmod +x /usr/local/bin/jenkinsapi
