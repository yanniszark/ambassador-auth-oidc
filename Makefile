REPO ?= gcr.io/arrikto-playground/istio-demo/ambassador-auth-oidc
VERSION ?= v0.3

all:
	docker build -t $(REPO):$(VERSION) .
	docker push $(REPO):$(VERSION)
