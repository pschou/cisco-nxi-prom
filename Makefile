PROG_NAME := "cisco-prom"
IMAGE_NAME := "pschou/${PROG_NAME}"
VERSION = 0.1.$(shell date +%Y%m%d.%H%M)
SRC = $(shell git config --get remote.origin.url | sed 's/.*@//;s/:/\//;s/\.git//' )
FLAGS := "-s -w -X main.version=${VERSION}@${SRC}"

build:
	CGO_ENABLED=0 go build -ldflags=${FLAGS} -o ${PROG_NAME}

linux:
	CGO_ENABLED=0 go build -ldflags=${FLAGS} -o ${PROG_NAME}
	upx --lzma ${PROG_NAME}
windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags=${FLAGS} -o ${PROG_NAME}.exe
	upx --lzma ${PROG_NAME}.exe

docker: linux
	DOCKER_BUILDKIT=1  docker image build -f Dockerfile --tag ${IMAGE_NAME}:${VERSION} .
	docker save ${IMAGE_NAME}:${VERSION} -o cisco-prom.tar

vendor:
	$(call mkvendor,github.com/pschou/go-cisco-nx-api)
	$(call mkvendor,github.com/pschou/go-json)
	$(call mkvendor,github.com/sirupsen/logrus)

.PHONY: vendor

define mkvendor
mkdir -p vendor/$(1)
rsync --prune-empty-dirs -a --exclude=.git --include=*/ --exclude=*_test.go --exclude=*_example.go --include=*.go --exclude=*  ~/go/src/$(1)/ vendor/$(1)
endef
