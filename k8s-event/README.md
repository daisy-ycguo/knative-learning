# Kubernetes Event Source Example

Kubernetes Event Source example shows how to wire kubernetes cluster events for consumption by a function that has been implemented as a Knative Service.

## Deploy
```
kubectl apply -f .
```
## 2. Create a in-memory-channel channel
## 3. Create a Service Account to subscribe K8s events
## 4. Create an event sourcing under default namespace
## 5. Create a subscription
## 6. Debugging