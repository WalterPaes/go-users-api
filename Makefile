mock-all:
	docker run -v "$PWD":/src -w /src vektra/mockery --all