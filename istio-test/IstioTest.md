# Testing istio on Knative

## Prepare Knative source code


## Install Knative
```
ko apply -f config/
```

## Test Knative is installed
```
kubectl -n knative-serving get pods
```

## Enable monitoring
```
kubectl apply -R -f config/monitoring/100-namespace.yaml \
    -f third_party/config/monitoring/logging/elasticsearch \
    -f config/monitoring/logging/elasticsearch \
    -f third_party/config/monitoring/metrics/prometheus \
    -f config/monitoring/metrics/prometheus \
    -f config/monitoring/tracing/zipkin
```

## Run helloworld
```
$ kubectl apply -f /Users/Daisy/workspace/knative-learning/service.yaml
$ export KINGRESS_IP=`kubectl get svc istio-ingressgateway --namespace istio-system --output jsonpath="{.status.loadBalancer.ingress[*].ip}"`
$ curl -H "Host: helloworld-go.default.example.com" http://$KINGRESS_IP

```

## Check data
Visit [Zipkin](http://localhost:8001/api/v1/namespaces/istio-system/services/zipkin:9411/proxy/zipkin/) to get trace after running:

```
kubectl proxy
```
