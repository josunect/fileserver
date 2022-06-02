# fileserver
http server that lists directory files
written in GO 

It is using a module not public, file manager, with go modules. 
- git clone git@github.com:josunect/fileserver.git
- cd filemanager
- go build
- cd ../httpserver
- go build 
- go run .


## Usage: 
go run .

Go to https://localhost:8443

## Routes:

- / 
- /list
- /directory

## Notes 
- Go Modules: https://go.dev/doc/tutorial/call-module-code 

## Certificate
- http server uses a localhost.crt and a localhot.key file in the httpserver folder
- Generated with:
  - openssl req  -new  -newkey rsa:2048  -nodes  -keyout localhost.key  -out localhost.csr
  - openssl  x509  -req  -days 365  -in localhost.csr  -signkey localhost.key  -out localhost.crt 

## Podman Imagen
- To build an image:
podman build -t httpserver .
- To run the image: 
podman run -p 8443:8443 --name httpserver localhost/httpserver

