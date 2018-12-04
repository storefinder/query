TAG?=latest
REGISTRY?=gcr.io/tmogoserverless

build:
	docker build -t storefinder/query .

push:
	docker tag storefinder/query:$(TAG) $(REGISTRY)/storefinder/query:$(TAG)
	docker push $(REGISTRY)/storefinder/query:$(TAG)