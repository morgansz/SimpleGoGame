# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
FROM golang:1.17 as builder

# Copy local code to the container image.
WORKDIR /go/src/app
COPY . .

# Build the command inside the container.
RUN go get -d -v ./...
RUN go install -v ./...

# Use a minimal Docker image to run the app
FROM gcr.io/distroless/base-debian10
COPY --from=builder /go/bin/app /app

# Run the web service on container startup.
CMD ["/app"]
