doc:
	@docker buildx build \
		--file docs/Dockerfile \
		--platform linux/amd64 \
		--load --tag appnest \
		docs/
# Jerry Rig
	@docker run --rm \
		--volume $(shell pwd)/docs:/home/node/app \
		--workdir /home/node/app appnest
	@cp docs/README.md .
	@rm docs/README.md
