# TODO build for linux

BINARY_NAME := app
CURRENT_DIR := ./
RUN_BINARY_CMD := $(CURRENT_DIR)$(BINARY_NAME)
PROJECT_NAME := hrubos.dev/collectorsden
DOC_FILENAME := generated

DOC_PATH_TMP = doc/$(DOC_FILENAME).tmp
DOC_PATH_MD = doc\$(DOC_FILENAME).md

MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

ICON_PATH = assets/img/icon_v3.png
ICON_PATH_ABS = $(MAKEFILE_DIR)$(ICON_PATH)

UNAME_S := $(OS)

FIREFOX := firefox
RMRF = rm -rf
RM = rm

# Detect OS and set remove command, doc path, and open command
ifeq ($(UNAME_S),Windows_NT)
    RM = del
	RMRF = del
    DOC_PATH_TMP = doc\$(DOC_FILENAME).tmp
	DOC_PATH_MD = doc\$(DOC_FILENAME).md
	FIREFOX = cmd /C start firefox
	EXE_STRING = .exe
	BINARY_NAME = $(BINARY_NAME)$(EXE_STRING)
	RUN_BINARY_CMD = $(BINARY_NAME)
endif

all: run

# with logs
build_debug:
	go build -ldflags="-X $(PROJECT_NAME)/internal/logger.debugBuildValue=true" -o $(BINARY_NAME) ./cmd

# without logs
build_release:
	go build -o $(BINARY_NAME) ./cmd

run: build_debug
	$(RUN_BINARY_CMD)

webserver:
	fyne serve -src cmd/ -icon "$(ICON_PATH_ABS)" --debug

doc:
	gomarkdoc --output doc/generated.tmp ./...
	go run scripts/stripheader.go doc/$(DOC_FILENAME).tmp doc/$(DOC_FILENAME).md
	$(RMRF) $(DOC_PATH_TMP)
	$(FIREFOX) doc\$(DOC_FILENAME).md

clean:
	go clean
	go mod tidy
	${RM} $(BINARY_NAME)

.PHONY: doc
.PHONY: clean
