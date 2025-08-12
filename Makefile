# TODO build for linux

BINARY_NAME := app.exe
PROJECT_NAME := hrubos.dev/collectorsden
DOC_FILENAME := generated

DOC_PATH_TMP = doc/$(DOC_FILENAME).tmp
DOC_PATH_MD = doc\$(DOC_FILENAME).md

UNAME_S := $(OS)

FIREFOX := firefox
RM = rm -rf

# Detect OS and set remove command, doc path, and open command
ifeq ($(UNAME_S),Windows_NT)
    RM = del
    DOC_PATH_TMP = doc\$(DOC_FILENAME).tmp
	DOC_PATH_MD = doc\$(DOC_FILENAME).md
	FIREFOX = cmd /C start firefox
endif

all: run

# with logs
build_debug:
	go build -ldflags="-X $(PROJECT_NAME)/internal/logger.debugBuildValue=true" -o $(BINARY_NAME) ./cmd

# without logs
build_release:
	go build -o $(BINARY_NAME) ./cmd

run: build_debug
	$(BINARY_NAME)

doc:
	gomarkdoc --output doc/generated.tmp ./...
	go run scripts/stripheader.go doc/$(DOC_FILENAME).tmp doc/$(DOC_FILENAME).md
	$(RM) $(DOC_PATH_TMP)
	$(FIREFOX) doc\$(DOC_FILENAME).md

clean:
	go clean
	go mod tidy
	del $(BINARY_NAME)

.PHONY: doc
.PHONY: clean
