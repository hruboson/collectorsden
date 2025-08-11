# TODO build for linux

BINARY_NAME = app.exe
PROJECT_NAME = hrubos.dev/collectorsden

all: run

# with logs
build_debug:
	go build -ldflags="-X ${PROJECT_NAME}/internal/logger.debugBuildValue=true" -o ${BINARY_NAME} ./cmd

# without logs
build_release:
	go build -o ${BINARY_NAME} ./cmd

run: build_debug
	${BINARY_NAME}

clean:
	go clean
	go mod tidy
	del ${BINARY_NAME}
