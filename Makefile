


build:
	CGO_ENABLED=0 go build -o cli client/main.go
	CGO_ENABLED=0 go build -o srv server/main.go
	docker build -t cast .



run:
	docker-compose up