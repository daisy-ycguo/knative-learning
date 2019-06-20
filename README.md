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

This repo contains useful steps to install Knative, test service on Knative and understand Knative.

## Install
1. Install istio
```
kubectl apply -f istio.yaml
kubectl label namespace default istio-injection=enabled
# Check pods are running
kubectl get pods -n istio-system
```
2. Install build and serving
```
kubectl apply -f serving-release-0.2.2.yaml
# Check pods are running
kubectl get pods -n knative-serving
kubectl get pods -n knative-build
```
3. Install your first app
```
kubectl apply -f service.yaml
# Check your pod is running
kubectl get Pods

# Get the public IP of knative-ingressgateway
export KINGRESS_IP=`kubectl get svc knative-ingressgateway --namespace istio-system --output jsonpath="{.status.loadBalancer.ingress[*].ip}"`
# Get url
export HELLOWORLD_URL=$(kubectl get services.serving.knative.dev helloworld-go  -o jsonpath='{.status.domain}')
# Test your app
curl -H "Host: ${HELLOWORLD_URL}" http://${KINGRESS_IP}

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

5. Eventing trial
```
#Install
kubectl apply -f eventing-release-0.2.0.yaml
#Install three eventsource type
kubectl apply -f eventing-eventsource-0.2.0.yaml

#Github samples
#Create a github-message-dumper service
kubectl apply -f ./github-event-sample/service.yaml
kubectl apply -f ./github-event-sample/githubsecret.yaml
kubectl apply -f ./github-event-sample/github-source.yaml
```

6. Logging, monitoring and metrics
#### Set up
**Note** If you are using IKS, add below mountPath and hostPath to DaemonSet `fluentd-ds`
```
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: fluentd-ds
  namespace: knative-monitoring
spec:
  template:
    spec:
      containers:
        volumeMounts:
        - mountPath: /var/data/cripersistentstorage
          name: persistentstorage
          readOnly: true
      volumes:
      - hostPath:
          path: /var/data/cripersistentstorage
        name: persistentstorage
```
Then
```
# Label node to start Fluentd DaemonSet
kubectl label nodes --all beta.kubernetes.io/fluentd-ds-ready="true"
# check
kubectl get daemonset fluentd-ds --namespace knative-monitoring
# start proxy
kubectl proxy
# start a local proxy of Grafana
kubectl port-forward --namespace knative-monitoring $(kubectl get pods --namespace knative-monitoring --selector=app=grafana --output=jsonpath="{.items..metadata.name}") 3000
```

#### Sample app: telemetry-go
Use `telemetry-go` as the sample app.
```
# Build docker image
cd $GOPATH/src/github.com/knative/docs
docker build \
  --tag "daisyycguo/knative-telemetry" \
  --file=serving/samples/telemetry-go/Dockerfile .
docker push daisyycguo/knative-telemetry

# Apply
kubectl apply -f ./telemetry-go/
export KINGRESS_IP=`kubectl get svc knative-ingressgateway --namespace istio-system --output jsonpath="{.status.loadBalancer.ingress[*].ip}"`
export TELEMETRY_HOST=`kubectl get route telemetrysample-route --output jsonpath="{.status.domain}"`
curl -H "Host:$TELEMETRY_HOST" http://${KINGRESS_IP}
curl -H "Host:$TELEMETRY_HOST" http://${KINGRESS_IP}/log

# Exporting
kubectl proxy
kubectl -n knative-monitoring port-forward $(kubectl -n knative-monitoring get pod -l app=prometheus -o jsonpath="{.items[0].metadata.name}") 9090
kubectl port-forward --namespace knative-monitoring $(kubectl get pods --namespace knative-monitoring --selector=app=grafana --output=jsonpath="{.items..metadata.name}") 3000
```
+ Visit [Kibana UI](http://localhost:8001/api/v1/namespaces/knative-monitoring/services/kibana-logging/proxy/app/kibana) to get logs. Use `kubernetes.labels.serving_knative_dev\/revision: : "telemetrysample-configuration-00001"` to search logs of this revision. Use `kubernetes.labels.serving_knative_dev\/configuration: "telemetrysample-configuration"` to search logs of this configuration.
+ Visit [Zipkin](http://localhost:8001/api/v1/namespaces/istio-system/services/zipkin:9411/proxy/zipkin/) to get trace.
+ Visit [Prometheus UI](http://localhost:9090) to get metrics. Use `istio_revision_request_duration_sum{destination_configuration="telemetrysample-configuration"}` and `istio_revision_request_count{destination_configuration="telemetrysample-configuration"}` to search metrics.
+ Visit [Grafana](http://localhost:3000) to get metrics.

```
# Customize prometheus
kubectl port-forward $(kubectl get pods --selector=app=prometheus-test --output=jsonpath="{.items[0].metadata.name}") 9090
# Access http://localhost:9090
```

Delete resources
```
kubectl delete -f ./telemetry-go/
```

### Sample app: blue/green deployment
```
kubectl apply -f ./blue-green-demo/blue-green-demo-config.yaml
kubectl apply -f ./blue-green-demo/blue-green-demo-route.yaml
export KINGRESS_IP=`kubectl get svc knative-ingressgateway --namespace istio-system --output jsonpath="{.status.loadBalancer.ingress[*].ip}"`
export BLUEGREEN_DEMO_HOST=`kubectl get route blue-green-demo --output jsonpath="{.status.domain}"`
curl -H "Host: $BLUEGREEN_DEMO_HOST" http://$KINGRESS_IP

# Change config from blue to green
# Change route to send traffic to V1 and send 0 traffic to V2 but with name route
kubectl apply -f ./blue-green-demo/blue-green-demo-config.yaml
kubectl apply -f ./blue-green-demo/blue-green-demo-route.yaml
curl -H "Host: $BLUEGREEN_DEMO_HOST" http://$KINGRESS_IP
curl -H "Host: v2.$BLUEGREEN_DEMO_HOST" http://$KINGRESS_IP

# Change route to half and half
kubectl apply -f ./blue-green-demo/blue-green-demo-route.yaml
curl -H "Host: $BLUEGREEN_DEMO_HOST" http://$KINGRESS_IP

# delete
kubectl delete route blue-green-demo
kubectl delete configuration blue-green-demo
```

## Appendix
Open prometheus Web UI:
```
kubectl -n knative-monitoring port-forward $(kubectl -n knative-monitoring get pod -l app=prometheus -o jsonpath="{.items[0].metadata.name}") 9090
```
kubectl port-forward $(kubectl get pods --selector=app=prometheus-test --output=jsonpath="{.items[0].metadata.name}") 9090
