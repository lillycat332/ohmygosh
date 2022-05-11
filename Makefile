BUILD_DIR := bin

cat: |$(BUILD_DIR)
	@echo "Building omgsh..."
	@go build -o bin/omgsh ./cmd/

$(BUILD_DIR):
	@echo "Folder $(BUILD_DIR) does not exist, creating it..."
	mkdir -p $@

