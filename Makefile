CONTAINER_IMAGE?=shopee-notify
RELEASE?=latest
MAIN_PATH?=Dockerfile

APP?=goapp
GOOS?=linux
GOARCH?=amd64

run:
	go run cmd/main.go
rmi:
	docker rmi $$(docker images | grep '$(CONTAINER_IMAGE)') || true
build-image: rmi
	docker build --no-cache -t $(CONTAINER_IMAGE):$(RELEASE) -f $(MAIN_PATH) .
up:
	docker-compose up -d
down:
	docker-compose down