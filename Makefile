pro:
	docker rmi -f web-service:1.0
	docker-compose up -d

dev:
	cd cmd/dev; go run main.go
	