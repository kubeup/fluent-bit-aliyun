DOCKER_IMAGE_NAME       ?= archon-controller
DOCKER_IMAGE_TAG        ?= $(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))

all: glide docker

glide:
	glide install

clean:
	rm -rf *.so *.h *~

docker:
	@echo ">> building docker image"
	@docker build -t "$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)" .
