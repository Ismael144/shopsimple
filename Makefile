PROTO_DIR := ./proto  
ENVOY_DIR := ./envoy
ENVOY_COMPOSE_LABEL := envoy
ENVOY_DESCRIPTOR := $(ENVOY_DIR)/descriptor.pb

# Envs 
DOCKER_TAG ?= latest
GO_VERSION ?= 1.25

.PHONY: bufpush descriptor envoy build up clean  

bufpush:  
	cd $(PROTO_DIR) && buf push 
descriptor: bufpush
	cd $(PROTO_DIR) && buf build -o ../$(ENVOY_DESCRIPTOR)
envoy: descriptor 
	docker compose build $(ENVOY_COMPOSE_LABEL) && docker compose up $(ENVOY_COMPOSE_LABEL) -d 
build: 
	@echo "Building docker images with tag: $(DOCKER_TAG)"
	docker compose build --build-arg GIT_COMMIT=$(GIT_HASH)
up: 
	docker compose up -d 
clean: 
	rm -f $(ENVOY_DESCRIPTOR)
