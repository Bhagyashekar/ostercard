PATH := $(GOPATH)/bin:$(PATH)
APP_EXECUTABLE="out/ostercard"

.PHONY: all
all: tidy build test

.PHONY: build
build:
	mkdir -p out/
	GO111MODULE=on CGO_ENABLED=0 go build -a -o $(APP_EXECUTABLE) ./cmd/ostercard

.PHONY: start
start:
	$(APP_EXECUTABLE)

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: test
test:
	go test -count=1 -covermode=atomic -coverprofile=coverage.out ./...
	go tool cover -html coverage.out -o coverage.html
	@coverage="$$(go tool cover -func coverage.out | grep 'total:' | awk '{print int($$3)}')"; \
	echo "The overall coverage is $$coverage%. Look at coverage.html for details."; \
	if [ $$coverage -lt 80 ]; then \
		echo "The coverage $$coverage% is below the accepted threshold 70%."; \
		exit 1; \
	fi
