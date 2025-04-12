start: build
	@docker compose up --build
build:
	-@docker compose rm -v --force --stop
	-@docker rmi ticket-booking-backend
	@docker compose build