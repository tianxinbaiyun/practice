DOCKER_NAME          	:= dockertest
VERSION					:= v1.0

dev:
	@$(info build and upload dev)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags "-s -w" -o server
	@rm -f $(DOCKER_NAME).img
	@docker rmi -f $(DOCKER_NAME):$(VERSION)
	docker build -t $(DOCKER_NAME):$(VERSION) .
	docker save -o $(DOCKER_NAME).img $(DOCKER_NAME)


tidy:
	GOPRIVATE=github.com/xs23933,e.coding.net/xs23933,git.huixing.tech;go mod tidy
