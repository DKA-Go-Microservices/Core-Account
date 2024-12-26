# Set the root project path and release folder path

ROOT_DIR := $(shell pwd)
RELEASE_DIR := $(ROOT_DIR)/release
TEST_DIR := $(ROOT_DIR)/test
EXECUTABLE_NAME := main # Define the desired executable name

build: clean proto
	@echo "TASK: Creating release directory ..."
	@mkdir -p $(RELEASE_DIR)
	@echo "TASK: Building ..."
	@cd cmd && go build -o $(RELEASE_DIR)/$(EXECUTABLE_NAME)

# Clean built binaries
clean:
	@echo "TASK: Cleaning up..."
	@rm -rf $(RELEASE_DIR)/*
	@rm -rf generated/*

# Build All Proto
proto: clean
	@echo "TASK: Compile All Proto"
	@mkdir -p generated
	@find api/ -type f -name "*.proto" | xargs protoc --go_out=./generated --go-grpc_out=./generated
	@find generated/ -type f -name "*.pb.go" | xargs -I {} protoc-go-inject-tag --input={}

# run dev
run: proto
	@echo "TASK: dev running app"
	@cd cmd && go run -v main.go

preview: build
	@echo "TASK: Run Executable Preview"
	@cd $(RELEASE_DIR) && ./$(EXECUTABLE_NAME)

docker: build
	@echo "TASK: Running Docker Compose"
	@docker compose up -d --force-recreate

prod:
	@echo "TASK: Run Production Executable"
	@cd $(RELEASE_DIR) && ./$(EXECUTABLE_NAME)

dev:
	@echo "TASK: Running Test Client GRPC"
	@cd $(TEST_DIR) && go run main.go

