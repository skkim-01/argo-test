TAG ?= v0.7.0
REGISTRY ?= skkim01/ubuntu-go
CONTAINER_NAME = ubuntu-go

docker-build:
	docker build -t ${REGISTRY}:${TAG} -f Dockerfile .

docker-push:
	docker push ${REGISTRY}:${TAG}

#
#
#
container-run:
	docker container run -d -t --name ${CONTAINER_NAME} ${REGISTRY}:${TAG}

container-stop:
	docker container stop ${CONTAINER_NAME}

container-rm:
	docker container rm ${CONTAINER_NAME}

exec:
	sudo docker exec -it ${CONTAINER_NAME} /bin/bash