start:
	go run cmd/main.go
build-go:
	GOOS=linux CGO_ENABLED=0  GOARCH=amd64 go build -o go-micro-proto-example cmd/main.go
init:
	gvm use go1.17.10
dashboard:
	docker run -d --name micro-dashboard -p 8082:8082 xpunch/go-micro-dashboard:latest
build-image:
	docker build -t nhddev/go-micro-proto-example .
publish-image:
	docker push nhddev/go-micro-proto-example:latest