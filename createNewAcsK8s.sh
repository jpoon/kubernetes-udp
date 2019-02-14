#!/bin/bash

set -ex

. var.env

## -------
## create resource group
az group create --name=$RESOURCE_GROUP --location=$LOCATION

## -------
## create kubernetes cluster
az aks create --resource-group $RESOURCE_GROUP --name=$CLUSTER_NAME --dns-prefix=$DNS_PREFIX
