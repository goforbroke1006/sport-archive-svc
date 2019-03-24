BUILD_PATH=./build/Realise
SERVICE_NAME=sport-archive-svc
ENTRY_POINT=cmd/${SERVICE_NAME}/main.go

DOCKER_IMAGE_NAME=goforbroke1006/sport-archive-svc

all: build

build: build-linux-amd64 build-linux-386

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o ${BUILD_PATH}/${SERVICE_NAME} ${ENTRY_POINT}

build-linux-386:
	GOOS=linux GOARCH=386   go build -o ${BUILD_PATH}/${SERVICE_NAME} ${ENTRY_POINT}

docker: build-linux-amd64 docker-build docker-push

docker-build:
	docker build --no-cache --build-arg BINARY_LOCATION=${BUILD_PATH}/${SERVICE_NAME} --build-arg BINARY_NAME=${SERVICE_NAME} -f docker/app/Dockerfile -t ${DOCKER_IMAGE_NAME}:latest .

docker-push:
	docker login
	docker push ${DOCKER_IMAGE_NAME}:latest
