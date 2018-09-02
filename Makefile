APPNAME = djreed/davidjreed.net

default: run

go:
	CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -a -installsuffix cgo -o ${APPNAME} .

docker:
	docker build -t ${APPNAME} -f Dockerfile .

run: docker
	docker run --rm -it -p 8080:8080 ${APPNAME}