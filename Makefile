default: fmt vet test build

BIN_FOLDER=bin
BIN_NAME=app

archive:
	tar czvf ${BIN_FOLDER}/${BIN_NAME}.tar.gz ${BIN_FOLDER}/amd64

clean:
	rm -rf ${BIN_FOLDER}

# Build on current system
build:
	go build -o ${BIN_FOLDER}/${BIN_NAME}

build-all: clean build-macos build-windows build-linux build-alpine-scratch

build-macos:
	GOOS=darwin GOARCH=amd64 go build -o ${BIN_FOLDER}/amd64/darwin/${BIN_NAME}

build-windows:
	GOOS=windows GOARCH=amd64 go build -o ${BIN_FOLDER}/amd64/windows/${BIN_NAME}.exe

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ${BIN_FOLDER}/amd64/linux/${BIN_NAME}

# Alpine & scratch base images use musl instead of gnu libc, thus we need to add additional parameters on the build
build-alpine-scratch:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ${BIN_FOLDER}/amd64/scratch/${BIN_NAME}

fmt:
	go fmt ./...

get:
	go get

# Generate binaries and an archive containing all binaries in bin/ folder
release: clean build-all archive

test:
	go test -v -timeout 60s -race ./...

# Colorized test results in terminal
test-colorized:
	@if type "richgo" > /dev/null 2>&1; then richgo test -v -timeout 60s -race ./...; else GO111MODULE=off go get github.com/kyoh86/richgo && richgo test -v -timeout 60s -race ./...; fi

vet:
	go vet ./...

watch:
	@if type "inotifywait" > /dev/null 2>&1; then while inotifywait -e close_write **/*.go -e close_write *.go; do make fmt vet test-colorized build; done; else echo 'Please install https://github.com/rvoicilas/inotify-tools/wiki'; fi
