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

# Knative Insight

This document describe the insights of Knative

## Describe a service pod
A service pod contains 1 init container `docker.io/istio/proxy_init` and 4 containers:
- user-container: the container with user codes;
- queue-proxy: Knative Serving owned sidecar container to enforce request concurrency limits, image of `gcr.io/knative-releases/github.com/knative/serving/cmd/queue@sha256`;
- fluentd-proxy: Sidecar container to collect logs from /var/log, image `k8s.gcr.io/fluentd-elasticsearch`;
- istio-proxy: Sidecar container to form an Istio mesh, image `docker.io/istio/proxyv2:1.0.2`, running `pilot-agent` and `envoy`.

## Two kinds of `knative.dev/type`
- app: long running service

- function: function as a service

##

helloworld-go
```
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: 2018-12-06T08:49:29Z
  name: helloworld-go
  namespace: default
  ownerReferences:
  - apiVersion: serving.knative.dev/v1alpha1
    blockOwnerDeletion: true
    controller: true
    kind: Route
    name: helloworld-go
    uid: d24d7f13-f933-11e8-8d79-062cfe122831
  resourceVersion: "1206457"
  selfLink: /api/v1/namespaces/default/services/helloworld-go
  uid: d7d58441-f933-11e8-8d79-062cfe122831
spec:
  externalName: knative-ingressgateway.istio-system.svc.cluster.local
  sessionAffinity: None
  type: ExternalName
status:
  loadBalancer: {}
```
telemetrysample-route
```
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: 2018-12-11T03:01:39Z
  name: telemetrysample-route
  namespace: default
  ownerReferences:
  - apiVersion: serving.knative.dev/v1alpha1
    blockOwnerDeletion: true
    controller: true
    kind: Route
    name: telemetrysample-route
    uid: bd5229ad-fcf0-11e8-97dc-86906447daf8
  resourceVersion: "1986783"
  selfLink: /api/v1/namespaces/default/services/telemetrysample-route
  uid: 14672b0a-fcf1-11e8-8d79-062cfe122831
spec:
  externalName: knative-ingressgateway.istio-system.svc.cluster.local
  sessionAffinity: None
  type: ExternalName
status:
  loadBalancer: {}
```
helloworld-go-00001-service
```
apiVersion: v1
kind: Service
metadata:
  annotations:
    serving.knative.dev/configurationGeneration: "1"
  creationTimestamp: 2018-12-06T08:49:22Z
  labels:
    app: helloworld-go-00001
    autoscaling.knative.dev/kpa: helloworld-go-00001
    serving.knative.dev/configuration: helloworld-go
    serving.knative.dev/revision: helloworld-go-00001
    serving.knative.dev/revisionUID: d2501208-f933-11e8-8d79-062cfe122831
    serving.knative.dev/service: helloworld-go
  name: helloworld-go-00001-service
  namespace: default
  ownerReferences:
  - apiVersion: serving.knative.dev/v1alpha1
    blockOwnerDeletion: true
    controller: true
    kind: Revision
    name: helloworld-go-00001
    uid: d2501208-f933-11e8-8d79-062cfe122831
  resourceVersion: "1206386"
  selfLink: /api/v1/namespaces/default/services/helloworld-go-00001-service
  uid: d3be7940-f933-11e8-8d79-062cfe122831
spec:
  clusterIP: 172.21.5.218
  ports:
  - name: http
    port: 80
    protocol: TCP
    targetPort: queue-port
  selector:
    serving.knative.dev/revision: helloworld-go-00001
  sessionAffinity: None
  type: ClusterIP
status:
  loadBalancer: {}
```
telemetrysample-configuration-00001-service
```
apiVersion: v1
kind: Service
metadata:
  annotations:
    serving.knative.dev/configurationGeneration: "1"
  creationTimestamp: 2018-12-11T03:01:24Z
  labels:
    app: telemetrysample-configuration-00001
    autoscaling.knative.dev/kpa: telemetrysample-configuration-00001
    knative.dev/type: app
    serving.knative.dev/configuration: telemetrysample-configuration
    serving.knative.dev/revision: telemetrysample-configuration-00001
    serving.knative.dev/revisionUID: bda484b3-fcf0-11e8-8d79-062cfe122831
    serving.knative.dev/service: ""
  name: telemetrysample-configuration-00001-service
  namespace: default
  ownerReferences:
  - apiVersion: serving.knative.dev/v1alpha1
    blockOwnerDeletion: true
    controller: true
    kind: Revision
    name: telemetrysample-configuration-00001
    uid: bda484b3-fcf0-11e8-8d79-062cfe122831
  resourceVersion: "1986692"
  selfLink: /api/v1/namespaces/default/services/telemetrysample-configuration-00001-service
  uid: 0b92b42e-fcf1-11e8-8d79-062cfe122831
spec:
  clusterIP: 172.21.183.183
  ports:
  - name: http
    port: 80
    protocol: TCP
    targetPort: queue-port
  selector:
    serving.knative.dev/revision: telemetrysample-configuration-00001
  sessionAffinity: None
  type: ClusterIP
status:
  loadBalancer: {}
```
