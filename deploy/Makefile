COMMIT_SHA=$(shell git rev-parse --short HEAD)
BUILD_DATE=$(shell date '+%Y-%m-%d %H:%M:%S %Z')
GO_VERSION=$(shell go version)
LD_FLAGS="-X 'main.hash=$(COMMIT_SHA)' -X 'main.builddate=$(BUILD_DATE)' -X 'main.goversion=$(GO_VERSION)'"

.PHONY: build

build:
	@echo "Building..."
	@cd ../src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags ${LD_FLAGS} -o ../deploy/build/go-http-server-template ./cmd/serve/main.go

TAG=$(subst /,_,$(shell git branch --show-current))
REPOSITORY_NAME=166812248980.dkr.ecr.ap-northeast-1.amazonaws.com/go-http-server-template
IMAGE_NAME=${REPOSITORY_NAME}:${TAG}
image: build
	@echo "Building image"
	docker build -t ${IMAGE_NAME} -f Dockerfile_for_push .

PROFILE=haruto-coating-plus-admin
ECR_URI=166812248980.dkr.ecr.ap-northeast-1.amazonaws.com
push-image: image
	@echo "Login AWS ECR..."
	@aws ecr get-login-password --region ap-northeast-1 --profile=${PROFILE} | docker login --username AWS --password-stdin ${ECR_URI}
	@echo "Pushing Image ..."
	@docker push ${IMAGE_NAME}
	@echo "Removing Local Image ..."
	@docker rmi ${IMAGE_NAME}
