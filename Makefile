BUILD_PATH=./build/Realise
SERVICE_NAME=sport-archive-svc
ENTRY_POINT=cmd/sport-archive-server/main.go

all: build

build: build-linux-amd64 build-linux-386

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o ${BUILD_PATH}/${SERVICE_NAME} ${ENTRY_POINT}

build-linux-386:
	GOOS=linux GOARCH=386   go build -o ${BUILD_PATH}/${SERVICE_NAME} ${ENTRY_POINT}


