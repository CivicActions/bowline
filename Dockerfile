# Build image
FROM golang:1.11.3-alpine as builder

# Systemwide setup
ENV GO111MODULE=on
RUN apk add --no-cache git gcc musl-dev

WORKDIR /go/src/bowline

# Build (rarely changing) module cache
COPY go.mod go.sum ./
RUN go mod download

# Build code files
COPY pkg ./pkg
COPY cmd ./cmd
RUN ls -R
RUN cd cmd/bowline ; go install -v

# Bowline image
FROM alpine:3.8

# Install bowline
COPY --from=builder /go/bin/bowline /usr/bin/bowline

# Install latest version of Docker Compose
RUN apk add py-pip && pip install docker-compose

WORKDIR /src
ENTRYPOINT [ "/usr/bin/bowline" ]