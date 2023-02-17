docker_run:
	docker compose build
	docker compose up

docker_stop:
	docker compose down

migrate:
	goose postgres 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' down

