include .env
export

export PROJECT_ROOT=$(shell pwd)

up-weather-app:
	@docker compose up -d weather-app

down-weather-app:
	@docker compose down weather-app

run-weather-app:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/weather-app/main.go