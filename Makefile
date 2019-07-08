BUILD_PATH=./build/bin
SERVICE_NAME=sport-archive-svc
ENTRY_POINT=cmd/${SERVICE_NAME}/main.go

DOCKER_IMAGE_NAME=goforbroke1006/sport-archive-svc
VERSION=1.0.0

all: build

gen:
	protoc \
	    --proto_path=$(GOPATH)/src \
	    --proto_path=$(GOPATH)/src/github.com/gogo/protobuf/protobuf \
	    --proto_path=$(@D)/ \
	    --gogofast_out=plugins=grpc:$(@D) \
	    ./api/proto/v1/sport-archive-svc.proto

build: build-linux-amd64 build-linux-386

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o ${BUILD_PATH}/${SERVICE_NAME} ${ENTRY_POINT}

build-linux-386:
	GOOS=linux GOARCH=386   go build -o ${BUILD_PATH}/${SERVICE_NAME} ${ENTRY_POINT}

docker: build-linux-amd64 docker-build docker-push

docker-build:
	docker build -f docker/app/Dockerfile --cache-from=${DOCKER_IMAGE_NAME}:${VERSION} \
	    --build-arg BINARY_LOCATION=${BUILD_PATH}/${SERVICE_NAME} \
	    --build-arg BINARY_NAME=${SERVICE_NAME} \
	    -t ${DOCKER_IMAGE_NAME}:${VERSION} -t ${DOCKER_IMAGE_NAME}:latest .

docker-push:
	docker login
	docker push ${DOCKER_IMAGE_NAME}:${VERSION}
	docker push ${DOCKER_IMAGE_NAME}:latest
