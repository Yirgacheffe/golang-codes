#!/bin/sh

# Run IBM MQ in docker, mount self-signed cert
# Change password for your Own
docker run --volume ~/.ssh/ibmmq:/etc/mqm/pki/keys/ati 
           --name AutoInvoiceMGR 
           --env LICENSE=accept 
           --env MQ_QMGR_NAME=GDSInvoiceMGR 
           --publish 1415:1414 
           --publish 9444:9443 
           --detach 
           --env MQ_APP_PASSWORD=passw0rd ibmcom/mq:latest