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

CMD ["cd frontend", "npm run-script build"]
RUN ls -la
CMD ["cd .."]
RUN ls -la
COPY frontend/build /app/build

# Replace go.mod
COPY httpserver/go.build.mod ./go.mod

# Build the binary
RUN go build .

EXPOSE 8443
CMD ["./httpserver"]
