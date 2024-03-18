all: test run

test: sync_worspace
	go test ./core/application_service -v
	go test ./infrastructure/controllers -v

run: sync_worspace prepare_env
	docker-compose build
	docker-compose up

prepare_env:
	$(shell if ! [[ -f .env ]]; then cp env .env; fi;)

sync_worspace:
	go work sync
