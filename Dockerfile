FROM golang:1.10-alpine3.8 as builder
RUN apk update && apk add git
ADD *.go /go/src/ambassador-auth-oidc/
WORKDIR /go/src/ambassador-auth-oidc
ADD Gopkg.toml .
ADD Gopkg.lock .
RUN go get github.com/golang/dep/cmd/dep
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /go/bin/ambassador-auth-oidc

FROM debian:stable
LABEL org.label-schema.vcs-url="https://github.com/ajmyyra/ambassador-auth-oidc"
LABEL org.label-schema.version="1.3"
RUN apt-get update && apt-get install -y ca-certificates curl && rm -rf /var/cache/apk/*
RUN addgroup --system auth && adduser --system --group auth
USER auth
WORKDIR /app
COPY --from=builder /go/bin/ambassador-auth-oidc /app/
ENTRYPOINT [ "/app/ambassador-auth-oidc" ]
