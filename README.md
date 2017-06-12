# kubernetes-udp

Simple UDP workload to run on Kubernetes

## Background

**Update:** Fix has been checked into Kubernetes. We'll have support of exposing UDP on Azure coming in 1.7. Fix has been cherry-picked to 1.6.5 (https://github.com/kubernetes/kubernetes/pull/46174).

-- 

Azure CloudProvider does not support exposing UDP through the Azure Load Balancer (https://github.com/kubernetes/kubernetes/issues/43683).

Cole Mickens is working on a fix (https://github.com/colemickens/kubernetes/commit/84a7250aab89a197aaafa5a065e3d5093bc33ee0) and has pushed out a private build: 
docker.io/colemickens/hyperkube-amd64:cmpr-cpfix

To use, ssh into the k8s master and update the `controller-manager` manifest found in `/etc/kubernetes/manifests`.
