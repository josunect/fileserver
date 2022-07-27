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

COPY frontend frontend
RUN apk add --update nodejs npm
WORKDIR /app/frontend
RUN ls -la
RUN npm run-script build
WORKDIR /app
RUN mv frontend/build .
RUN rm -r frontend

# Replace go.mod
COPY httpserver/go.build.mod ./go.mod

# Build the binary
RUN go build .

EXPOSE 8443
CMD ["./httpserver"]
