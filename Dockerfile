FROM golang:alpine

# Copy core library
RUN mkdir /httpserver
COPY filemanager/ ./src/filemanager
WORKDIR ./src/filemanager
RUN go build
RUN ls -la

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY httpserver ./
COPY httpserver/go.build.mod ./go.mod
RUN ls -la
#RUN go mod download

RUN pwd
# Build the binary
RUN go build .
RUN ls -la

EXPOSE 8443
CMD ["./httpserver"]