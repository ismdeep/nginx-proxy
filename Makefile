help:

build-vendor:
	go mod tidy
	go mod vendor

docker-release:
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--pull \
		--push \
		-t hub.deepin.com/public/nginx-proxy:latest .
