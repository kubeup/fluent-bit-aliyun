DOCKER_IMAGE_NAME       ?= archon-controller
DOCKER_IMAGE_TAG        ?= $(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))

all: glide plugin

glide:
	glide install

plugin:
	go build -buildmode=c-shared -o out_sls.so .
	cp out_sls.so docker-image/

clean:
	rm -rf *.so *.h *~

docker:
	@echo ">> building docker image"
	@docker build -t "$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)" docker-image
