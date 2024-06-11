#!/bin/bash

kubectl apply -f .k8s/cluster-role.yaml

POD_NAME=$(kubectl get pods -A | grep rabbitmq-operator-controller | awk '{print $2}')

if [ -z "$POD_NAME" ]; then
  echo "Pod rabbitmq-operator-controller não encontrado."
  exit 1
fi

kubectl delete pod $POD_NAME -n rabbitmq-operator-system

if [ $? -eq 0 ]; then
  echo "Pod $POD_NAME deletado com sucesso."
else
  echo "Falha ao deletar o pod $POD_NAME."
  exit 1
fi

POD_NAME=$(kubectl get pods -A | grep rabbitmq-operator-controller | awk '{print $2}')

if [ -z "$POD_NAME" ]; then
  echo "Pod rabbitmq-operator-controller não encontrado."
  exit 1
fi

echo "Aguardando o pod $POD_NAME ficar pronto..."
sleep 20
echo "Logs do POD: $POD_NAME"

kubectl logs $POD_NAME -n rabbitmq-operator-system