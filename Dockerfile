# Build image
FROM golang:alpine as builder

# Get specific docker library version
RUN apk add --no-cache git && \
    git clone --single-branch -b v18.06.1-ce https://github.com/docker/engine.git $GOPATH/src/github.com/docker/docker

WORKDIR /go/src/bowline

COPY main.go .

RUN go get -v && go install -v


# Bowline image
FROM alpine:3.8

# Install bowline
COPY --from=builder /go/bin/bowline /usr/bin/bowline

# Install latest version of Docker Compose
RUN apk add py-pip && pip install docker-compose


WORKDIR /src
ENTRYPOINT /usr/bin/bowline
