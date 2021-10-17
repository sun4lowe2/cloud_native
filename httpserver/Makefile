export tag=v1.0
root:
	export ROOT=github.com/sun4lowe2/cloud_native

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpserver container"
	docker build -t 3un4lowe2/httpserver:${tag} .

push: release
	echo "pushing 3un4lowe2/httpserver"
	docker push 3un4lowe2/httpserver:${tag}
