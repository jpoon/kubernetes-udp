# kubernetes-udp

Simple UDP workload to run on Kubernetes

## Background

UDP works with AKS out of the box. Verified February 14th, 2019 with
```
$ kubectl version
Client Version: version.Info{Major:"1", Minor:"11", GitVersion:"v1.11.2", GitCommit:"bb9ffb1654d4a729bb4cec18ff088eacc153c239", GitTreeState:"clean", BuildDate:"2018-08-08T16:31:10Z", GoVersion:"go1.10.3", Compiler:"gc", Platform:"darwin/amd64"}
Server Version: version.Info{Major:"1", Minor:"9", GitVersion:"v1.9.11", GitCommit:"1bfeeb6f212135a22dc787b73e1980e5bccef13d", GitTreeState:"clean", BuildDate:"2018-09-28T21:35:22Z", GoVersion:"go1.9.3", Compiler:"gc", Platform:"linux/amd64"}
```

## Getting Started

- Create your Azure resources
```
az group create -l centralus -n xtoph-delete-aks-udp

az aks create -g xtoph-delete-aks-udp -n udpcluster

az aks get-credentials -g xtoph-delete-aks-udp -n udpcluster
```

- Create your containers
```
docker build -f ./server/Dockerfile -t xtoph/aks-udp-server .

docker build -f ./client/Dockerfile -t xtoph/aks-udp-client .

```

- Push containers to your Container Registry
```
docker push xtoph/aks-udp-server

docker push xtoph/aks-udp-client
```

- Deploy server and get IP
```
kubectl create -f server.yaml
```

- Get LoadBalancer Server IP address
```
kubectl get svc --selector=app=udp-server -o jsonpath='{.items[0].status.loadBalancer.ingress[0].ip}'
```

- Update `client.yaml`
```
    spec:
      containers:
      - name: udp-client
        image: jpoon/udp-client
        imagePullPolicy: Always
        env:
        - name: SERVER_ADDRESS
          # Replace with ClusterIP/ServiceIP
          value: "<LOAD BALANCER IP>"
```
## Test UDP

You can verify UDP connectivity through one of the following methods:

1. Running the container with the Client App from your local machine
```
$ docker run -p 10001 -it -e SERVER_ADDRESS=<LOAD BALANCER IP>  xtoph/aks-udp-client
SERVER_ADDRESS=<LOAD BALANCER IP>
ServerAddr=<LOAD BALANCER IP>:10001
Received  udp-server-deployment-654dc47f47-986h2  from  <LOAD BALANCER IP>:10001
Received  udp-server-deployment-654dc47f47-986h2  from  <LOAD BALANCER IP>:10001
```
1. Running the container with the Client App from the kubernetes cluster
```
$ kubectl create -f client.yaml
deployment.extensions/udp-client-deployment created
```

Then looking at the logs for the container you will see:
```
SERVER_ADDRESS=<LOAD BALANCER IP>
ServerAddr=<LOAD BALANCER IP>:10001
Received  udp-server-deployment-654dc47f47-72zf5  from  <LOAD BALANCER IP>:10001
Received  udp-server-deployment-654dc47f47-72zf5  from  <LOAD BALANCER IP>:10001
```
