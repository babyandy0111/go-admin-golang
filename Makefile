PROJECT:=go-admin

.PHONY: build
build:
	 make doc && env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./main && docker build -t go-admin:latest . --no-cache

dev:
	go run main.go server -c config/dev.settings.yml -a false

init:
	go run main.go migrate -c config/dev.settings.yml

doc:
	swag init --instanceName admin --parseDependency --parseInternal -o docs/admin