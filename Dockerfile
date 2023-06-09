# Use base golang image from Docker Hub
FROM golang:1.19 AS build

WORKDIR /app/go/base

# Install dependencies in go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the application source code
COPY . .

# Compile the application to /app.
# Skaffold passes in debug-oriented compiler flags
ARG SKAFFOLD_GO_GCFLAGS
RUN echo "Go gcflags: ${SKAFFOLD_GO_GCFLAGS}"
RUN CGO_ENABLED=0 go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -mod=readonly -v -o main

# Now create separate deployment image
FROM gcr.io/distroless/base:debug AS base

# Definition of this variable is used by 'skaffold debug' to identify a golang binary.
# Default behavior - a failure prints a stack trace for the current goroutine.
# See https://golang.org/pkg/runtime/
ENV GOTRACEBACK=single

# Copy template & assets
WORKDIR /app/go/base
COPY --from=build /app/go/base/main /app/go/src/main

ENTRYPOINT ["/app/go/src/main"]
