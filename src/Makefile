all: test run

test: sync_worspace
	go test ./core/application_service -v
	go test ./infrastructure/controllers -v

run: sync_worspace
	go run .

sync_worspace:
	go work sync
