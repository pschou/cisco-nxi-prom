PROG_NAME := "cisco-prom"
IMAGE_NAME := "pschou/${PROG_NAME}"
VERSION = 0.1.$(shell date +%Y%m%d.%H%M)
SRC = $(shell git config --get remote.origin.url | sed 's/.*@//;s/:/\//;s/\.git//' )
FLAGS := "-s -w -X main.version=${VERSION}@${SRC}"


build:
	CGO_ENABLED=0 go build -ldflags=${FLAGS} -o ${PROG_NAME}
	upx --lzma ${PROG_NAME}
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags=${FLAGS} -o ${PROG_NAME}.exe
	upx --lzma ${PROG_NAME}.exe

docker: build
	docker build -f Dockerfile --tag ${IMAGE_NAME}:${VERSION} .
	docker push ${IMAGE_NAME}:${VERSION};  
	docker save ${IMAGE_NAME}:${VERSION} > pschou_cisco-prom.tar
