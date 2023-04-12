.PHONY: cover
cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

.PHONY: gen
gen:
	mockgen -source=internal/storage/memstorage/mem_storage.go \
	-destination=internal/storage/memstorage/mock/mem_storage_mock.go
#	mockgen -source=internal/db/database.go \
#	-destination=internal/db/mock/database.go
#	mockgen -source=internal/config/configagent/config_agent.go \
#	-destination=internal/config/configagent/mock/config_agent.go
#	mockgen -source=internal/config/configserver/config_server.go \
#	-destination=internal/config/configserver/mock/config_server.go
	mockgen -source=internal/api/hashservicer/hash_servicer.go \
	-destination=internal/api/hashservicer/mock/hash_servicer_mock.go
#	mockgen -source=internal/api/server/handlers/handler.go \
#	-destination=internal/api/server/handlers/mock/handler_mock.go
#	mockgen -source=internal/api/file/file.go \
#	-destination=internal/api/file/mock/file.go

test:
	go test -v -count=1 ./...

test100:
	go test -v -count=100 ./...

race:
	go test -v -race -count=1 ./..
