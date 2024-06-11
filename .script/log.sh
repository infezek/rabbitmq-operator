#!/bin/bash

POD_NAME=$(kubectl get pods -A | grep rabbitmq-operator-controller | awk '{print $2}')
if [ -z "$POD_NAME" ]; then
  echo "Pod rabbitmq-operator-controller não encontrado."
  exit 1
fi

kubectl logs $POD_NAME -n rabbitmq-operator-system --tail=20 -f