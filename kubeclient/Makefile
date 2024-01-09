CONTAINER_NAME=kube-client

docker-build:
	docker build -t skkim-01/kubeclient:v1.0 -f Dockerfile .

container-run:
	docker container run -d -t -p 9999:9999 --name ${CONTAINER_NAME} skkim-01/kubeclient:v1.0

container-stop:
	docker container stop ${CONTAINER_NAME}

container-rm:
	docker container rm ${CONTAINER_NAME}

test-stat:
	curl -XGET http://127.0.0.1:8080/api/v1/stat

test-pods:
	curl -XGET http://127.0.0.1:8080/api/v1/cli/pods