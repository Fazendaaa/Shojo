build-docs:
	@docker buildx build \
		--file docs/Dockerfile \
		--platform linux/amd64 \
		--push --tag fazenda/appnest \
		docs/

doc:
	@docker run --rm \
		--volume ${PWD}:${PWD} \
		--workdir ${PWD} fazenda/appnest \
		--input docs/blueprint.md \
		--output README.md

builds:
	@docker buildx build \
		--file build/package/Dockerfile \
		--platform linux/amd64 \
		--load --tag shojo \
		.
	@docker buildx build \
		--platform linux/amd64 \
		--load --tag shojo-latex \
		.
