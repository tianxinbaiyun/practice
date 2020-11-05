DOCKER_NAME          := amazon
PPORT                := 10304
VERSION              := v2.0.1

dev:
	@$(info build and upload dev)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags "-s -w" -o server
	@rm -f $(DOCKER_NAME).img
	docker build -t $(DOCKER_NAME):$(VERSION) .
	docker save -o $(DOCKER_NAME).img $(DOCKER_NAME)
	docker rmi $(DOCKER_NAME):$(VERSION)
	rm -f server

	docker stop $(DOCKER_NAME) && docker rm $(DOCKER_NAME);\
	docker rmi $(DOCKER_NAME):$(VERSION);\
	docker load --input $(DOCKER_NAME).img;\
	docker run --restart=always \
		--ip 127.0.0.1 \
		-p 10304:10304 \
        --name $(DOCKER_NAME) -itd $(DOCKER_NAME):$(VERSION) ;\
	rm -f $(DOCKER_NAME).img;docker logs -f $(DOCKER_NAME)