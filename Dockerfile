FROM golang:alpine

# Copy core library
RUN mkdir /httpserver
COPY filemanager/ ./src/filemanager
WORKDIR ./src/filemanager
RUN go build

# Create and change to the app directory.
WORKDIR /app

# Move files
COPY httpserver ./
# Replace go.mod
COPY httpserver/go.build.mod ./go.mod

# Build the binary
RUN go build .

EXPOSE 6080
CMD ["./httpserver"]