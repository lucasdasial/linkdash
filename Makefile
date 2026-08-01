dev:
	source .env && go run .
	
setup:
	source .env && \
	docker compose down && \
	docker compose up -d && \
	sleep 1 && \
	goose up