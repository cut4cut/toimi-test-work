run: ### Run docker-compose
	docker-compose up --build -d app && docker-compose logs -f
.PHONY: run

down: ### Down docker-compose
	docker-compose down --remove-orphans
.PHONY: down

