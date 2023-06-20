start-server:
	go run cmd/main.go

init:
	gvm use go1.17.10

dashboard:
	docker run -d --name micro-dashboard -p 8082:8082 xpunch/go-micro-dashboard:latest