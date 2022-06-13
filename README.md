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
export CERTIFICATE_FILE="$HOME/data/httpserver/mycert.crt"
export CERTIFICATE_KEY="$HOME/data/httpserver/mycert.key"

go run .

Go to https://localhost:8443

or, to run in dev mode and serve as http:

go run . -env=dev 

## Routes:

- / 
- /list
- /directory

## Notes 
- Go Modules: https://go.dev/doc/tutorial/call-module-code

## Podman Image
- To build an image:
podman build -t httpserver .
- To run the image: 
podman run -p 8443:8443 --name httpserver localhost/httpserver

## Docker Image
- To build an image:
  docker build -t httpserver .

# Kubernetes
The example has been done for kubernetes and minikube
- minikube start --driver=docker

We need to map the file cert directory: (The image expects a mycert.crt and mycert.key)

- minikube mount $HOME/data:/usr/secrets

(So we don't need volumne.yaml)

- kubectl create -f configmap.yaml
- kubectl create -f deployment.yaml
- kubectl apply -f service.yaml
- kubectl apply -f ingress.yaml 

Access the service in the url provided. Access the url minikube ip

https://192.168.59.101/list

Note: The ingress.yaml is defined for Ingress controller ingress-nginx

# Image releases
0.5 Use confimap