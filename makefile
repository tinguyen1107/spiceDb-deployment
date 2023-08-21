start:
	@docker compose up -d
	sleep 10
	@sh setup.sh

clear:
	@docker compose down
	@docker volume rm demo_postgres_data

