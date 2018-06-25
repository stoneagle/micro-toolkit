PWD := $(shell pwd)
USER := $(shell id -u)
USERNAME := $(shell id -u -n)
GROUP := $(shell id -g)
PROJECT := toolkit
IDENTIFY_GIT_TAG := $(shell git describe --tags `git rev-list --tags --max-count=1`) 
DEVELOP_PREFIX = registry.cn-beijing.aliyuncs.com/roobodev/
RELEASE_PREFIX = registry.cn-beijing.aliyuncs.com/roobo/
GOVERSION = 1.10

run-web: 
	cd hack && docker-compose -p "$(PROJECT)-web-$(USER)" up
stop-web: 
	cd hack && docker-compose -p "$(PROJECT)-web-$(USER)" stop 
rm-web: 
	cd hack && docker-compose -p "$(PROJECT)-web-$(USER)" rm 

run-ng:
	cd frontend && ng serve --environment=dev

run-web-ol: rm-web
	export IDENTIFY_GIT_TAG=$(IDENTIFY_GIT_TAG) && \
	export IMG_PREFIX=$(DEVELOP_PREFIX) && \
	cd hack && docker-compose -f docker-compose-online.yaml -p "$(PROJECT)-web-$(USERNAME)-online" up
stop-web-ol: 
	export IDENTIFY_GIT_TAG=$(IDENTIFY_GIT_TAG) && \
	export IMG_PREFIX=$(DEVELOP_PREFIX) && \
	cd hack && docker-compose -f docker-compose-online.yaml -p "$(PROJECT)-web-$(USERNAME)-online" stop 
rm-web-ol: 
	export IDENTIFY_GIT_TAG=$(IDENTIFY_GIT_TAG) && \
	export IMG_PREFIX=$(DEVELOP_PREFIX) && \
	cd hack && docker-compose -f docker-compose-online.yaml -p "$(PROJECT)-web-$(USERNAME)-online" rm 

build-golang:
	cd hack && docker build -f ./Dockerfile -t toolkit/golang:1.10 .

build: clean-files stop-web rm-web stop-web-ol rm-web-ol ng-build backend-build tool-build build-service

build-service:
	docker build -f ./hack/release/Dockerfile -t $(DEVELOP_PREFIX)toolkit:$(IDENTIFY_GIT_TAG) .

push-service:
	docker push $(DEVELOP_PREFIX)toolkit:$(IDENTIFY_GIT_TAG)

clean-files:
	rm -rf ./backend/static/* && \
	rm -rf ./backend/release/* 

build-release:
	cp ./hack/release/Dockerfile.release.example ./hack/release/Dockerfile.release && \
	sed -i "s:IMAGE:$(DEVELOP_PREFIX)toolkit:" ./hack/release/Dockerfile.release && \
	sed -i "s:TAG:$(RELEASE_TAG):" ./hack/release/Dockerfile.release && \
	docker build -f ./hack/release/Dockerfile.release -t $(RELEASE_PREFIX)toolkit:$(RELEASE_TAG) . && \
	rm ./hack/release/Dockerfile.release

push-release:
	docker push $(RELEASE_PREFIX)toolkit:$(RELEASE_TAG)

backend-build:
	docker run -it --rm \
		-u $(USER):$(GROUP) \
		-v $(PWD)/release:/tmp/release \
		-v $(PWD)/backend:/go/src/toolkit/backend \
		-w /go/src/toolkit/backend \
		golang:$(GOVERSION) \
		go build -o /tmp/release/backend

tool-build:
	docker run -it --rm \
		-u $(USER):$(GROUP) \
		-v $(PWD)/release:/tmp/release \
		-v $(PWD)/backend:/go/src/toolkit/backend \
		-w /go/src/toolkit/backend/initial \
		golang:$(GOVERSION) \
		go build -o /tmp/release/tool

ng-build: 
	docker run -it --rm \
		-u $(USER):$(GROUP) \
		-v $(PWD):/app \
		-w /app/frontend \
		alexsuch/angular-cli:v1.1.3 \
		ng build --environment=prod --deploy-url static

# init
init-db:
	docker exec -w /go/src/toolkit/backend -it toolkit-$(USERNAME)-golang go run ./initial/init.go 
