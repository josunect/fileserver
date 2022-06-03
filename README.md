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
- eval $(minikube docker-env) # use local images created in the docker registry 
- minikube docker-env
- kubectl create -f deployment.yaml 
- kubectl get pods 
- kubectl describe deployment httpserver
- kubectl apply -f service.yaml
- kubectl get svc httpserver
- minikube service httpserver --url 
- kubectl apply -f ingress.yaml 
- kubectl get ingress
- Access the service in the url provided

Note: The ingress.yaml is defined for Ingress controller ingress-nginx