FROM golang:1.10.1 as builder

WORKDIR /go/src/github.com/mchmarny/knative-route-demo/
COPY . .

# restore to pinnned versions of dependancies 
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

# build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o knative-route-demo \
    -tags netgo -installsuffix netgo .

# build the clean image
FROM scratch as runner
# copy the app
COPY --from=builder /go/src/github.com/mchmarny/knative-route-demo/knative-route-demo .
# copy static artifacts 
COPY --from=builder /go/src/github.com/mchmarny/knative-route-demo/static /static
COPY --from=builder /go/src/github.com/mchmarny/knative-route-demo/templates /templates

ENTRYPOINT ["/knative-route-demo"]
