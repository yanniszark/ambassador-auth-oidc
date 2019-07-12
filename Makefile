IMAGE ?= "gcr.io/arrikto/kubeflow/oidc-authservice:v0.3"

docker-build:
	docker build -t $(IMAGE) .

docker-push:
	docker push $(IMAGE)

publish: docker-build docker-push
