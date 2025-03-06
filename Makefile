build:
	docker build -t bakeoff-app .

run:
	docker run -p 8081:8081 bakeoff-app

run-local:
	HUB_API_URL="http://hub.test:3000" go run ./cmd/bakeoff

.PHONY: run build run-local
