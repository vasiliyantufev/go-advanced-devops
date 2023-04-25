.PHONY: cover
cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

.PHONY: gen
gen:
	mockgen -source=internal/storage/memstorage/mem_storage.go \
	-destination=internal/storage/memstorage/mock/mem_storage_mock.go
	mockgen -source=internal/db/database.go \
	-destination=internal/db/mock/database.go
#	mockgen -source=internal/config/configagent/config_agent.go \
#	-destination=internal/config/configagent/mock/config_agent.go
#	mockgen -source=internal/config/configserver/config_server.go \
#	-destination=internal/config/configserver/mock/config_server.go
	mockgen -source=internal/api/hashservicer/hash_servicer.go \
	-destination=internal/api/hashservicer/mock/hash_servicer_mock.go
#	mockgen -source=internal/api/server/handlers/handler.go \
#	-destination=internal/api/server/handlers/mock/handler_mock.go
	mockgen -source=internal/storage/filestorage/file_storage.go \
	-destination=internal/storage/filestorage/mock/file_storage_mock.go

test:
	go test -v -count=1 ./...

test100:
	go test -v -count=100 ./...

race:
	go test -v -race -count=1 ./..

run_server:
	go run ./cmd/server/main.go

run_agent:
	go run ./cmd/agent/main.go

run_pprof_profile:
	go tool pprof -http=":9090" -seconds=30 http://localhost:8088/debug/pprof/profile

run_pprof_heap:
	go tool pprof -http=":9099" -seconds=30 http://localhost:8088/debug/pprof/heap

#docs url - http://localhost:6060/pkg/?m=all
run_godoc:
	godoc -http=:6060

run_multichecker:
	go run cmd/staticlint/multichecker.go ./...

gen_cripto_key:
	openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -out server.crt -subj "/CN=localhost" -addext "subjectAltName = DNS:localhost"