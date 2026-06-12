export PROJECT_ROOT=$(shell pwd)

delpoy-app:
	@docker-compose up -d test-app