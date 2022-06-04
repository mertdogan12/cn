FROM golang:latest as build-deps

# RUN apt-get update && apt-get install -y ca-certificates --no-install-recommends

WORKDIR /go/src/github.com/mertdogan12/cn

ENV GOPATH=/go

COPY . ./
RUN go get -d ./...
RUN go build .

# RUN rm -rf /root/.ssh

FROM ubuntu:latest

WORKDIR /app

# COPY --from=build-deps /etc/ssl/certs/ /etc/ssl/certs/
COPY --from=build-deps /go/src/github.com/mertdogan12/cn/cn ./

EXPOSE 3000

CMD ["/app/cn"]
