# query
Query service

## Build query service
At this point we should be able to build the query service. Since the build for query leverages a shareable knative CI build template be sure to create it before running the build as well as ensure service account credentials and service account resources is created in k8s

```
kubectl apply -f build.yaml
```

### Checking build
Run command below to check the status of build

```
kubectl get builds -o yaml
```
