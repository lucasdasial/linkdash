setup:
	source .env && \
	docker compose down && \
	docker compose up -d && \
	sleep 1 && \
	goose up

run: 
	go run main.go