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
RUN go test -v ./...

# Build minimal static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/bowline cmd/bowline/main.go

# Create Bowline image
FROM scratch

# Install bowline
COPY --from=builder /go/bin/bowline /go/bin/bowline

WORKDIR /src
ENTRYPOINT [ "/go/bin/bowline" ]
