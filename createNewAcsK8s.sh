#!/bin/bash

set -ex

NAME="japoon-situm"

## -------
## create service principal
SUBSCRIPTION_ID="04f7ec88-8e28-41ed-8537-5e17766001f5"
SERVICE_PRINCIPAL_NAME=$NAME
SERVICE_PRINCIPAL_PASSWORD=`date | md5 | head -c8; echo`
az ad sp create-for-rbac --name $SERVICE_PRINCIPAL_NAME --role="Contributor" --scopes="/subscriptions/$SUBSCRIPTION_ID" --password $SERVICE_PRINCIPAL_PASSWORD

## -------
## create resource group
RESOURCE_GROUP=$NAME
LOCATION=westus
az group create --name=$RESOURCE_GROUP --location=$LOCATION

## -------
## create kubernetes cluster
DNS_PREFIX=$NAME
CLUSTER_NAME=$NAME
az acs create --orchestrator-type=kubernetes --resource-group $RESOURCE_GROUP --name=$CLUSTER_NAME --dns-prefix=$DNS_PREFIX --service-principal http://$SERVICE_PRINCIPAL_NAME --client-secret $SERVICE_PRINCIPAL_PASSWORD
