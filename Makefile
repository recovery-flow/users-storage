OPENAPI_GENERATOR := java -jar ~/openapi-generator-cli.jar
CONFIG_FILE := ./config_local.yaml
API_SRC := ./docs/api.yaml
API_BUNDLED := ./docs/api-bundled.yaml
OUTPUT_DIR := ./docs/web
RESOURCES_DIR := ./resources

generate-models:
	find $(RESOURCES_DIR) -type f ! \( -name "resources_types.go" -o -name "links.go" \) -delete
	swagger-cli bundle $(API_SRC) --outfile $(API_BUNDLED) --type yaml

	$(OPENAPI_GENERATOR) generate \
		-i $(API_BUNDLED) -g go \
		-o $(OUTPUT_DIR) \
		--additional-properties=packageName=resources

	mkdir -p $(RESOURCES_DIR)
	find $(OUTPUT_DIR) -name '*.go' -exec mv {} $(RESOURCES_DIR)/ \;
	find $(RESOURCES_DIR) -type f -name "*_test.go" -delete

migrate-up:
	KV_VIPER_FILE=$(CONFIG_FILE) go build -o main main.go
	KV_VIPER_FILE=$(CONFIG_FILE) ./main migrate up

migrate-down:
	migrate -path internal/data/migration -database $(DB_URL) -verbose down

run-server:
	KV_VIPER_FILE=$(CONFIG_FILE) go build -o main ./cmd/users-storage/main.go
	KV_VIPER_FILE=$(CONFIG_FILE) ./main run service