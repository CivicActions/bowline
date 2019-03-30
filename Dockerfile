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
COPY fixtures ./fixtures
RUN go test -v ./... ; go install -v ./...

# Run golang ci lint
RUN wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.12.5
RUN golangci-lint run --enable-all ./...

# Bowline image
FROM docker/compose:1.24.0

# Install bowline
COPY --from=builder /go/bin/bowline /usr/bin/bowline

WORKDIR /src
ENTRYPOINT [ "/usr/bin/bowline" ]
