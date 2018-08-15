<!--
#
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
-->

# Study Knative

This repo contains useful steps to install Knative, test app on Knative and understand Knative.

## Install
1. Install istio
```
kubectl apply -f istio.yaml
# Check pods are running
kubectl get pods -n istio-system
```
2. Install build and serving
```
kubectl apply -f serving-release.yaml
# Check pods are running
kubectl get pods -n knative-serving
kubectl get pods -n knative-build
```
3. Install your first app
```
kubectl apply -f service.yaml
# Check your pod is running
kubectl get Pods

# Get ip address of istio-ingressgateway
kubectl get pods -n istio-system -o wide | grep "istio-ingressgateway"
# Get port
echo $(kubectl get svc knative-ingressgateway -n istio-system   -o 'jsonpath={.spec.ports[?(@.port==80)].nodePort}')
# Get the public IP for the private IP if you are using a public cloud
export IP_ADDRESS=<your public ip>:<port>
# Get url
export HOST_URL=$(kubectl get services.serving.knative.dev helloworld-go  -o jsonpath='{.status.domain}')
# Test your app
curl -H "Host: ${HOST_URL}" http://${IP_ADDRESS}

# Delete your deployment
kubectl delete -f service.yaml
```
4. A source to url sample
```
#Install the kaniko build template
kubectl apply -f kaniko.yaml
#Register secrets for Docker Hub
kubectl apply -f docker-secret.yaml
#Build and deploy
kubectl apply -f builder.yaml
# Get ip address of istio-ingressgateway
kubectl get pods -n istio-system -o wide | grep "istio-ingressgateway"
#Get the port
echo $(kubectl get svc knative-ingressgateway -n istio-system   -o 'jsonpath={.spec.ports[?(@.port==80)].nodePort}')
# Get the public IP for the private IP if you are using a public cloud
export IP_ADDRESS=<your public ip>:<port>
#Get the app URL
export HOST_URL=$(kubectl get services.serving.knative.dev app-from-source  -o jsonpath='{.status.domain}')
# Test your app
curl -H "Host: ${HOST_URL}" http://${IP_ADDRESS}

# Delete your deployment
kubectl delete -f builder.yaml
kubectl delete -f docker-secret.yaml
kubectl delete -f kaniko.yaml
```
