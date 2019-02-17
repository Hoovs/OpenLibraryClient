TEST_FLAGS := -v -race
vendor:
	dep ensure
build:
	mkdir build

build/OpenLibary: build
	go build -o build/OpenLibraryServer ./server

cleanBuild:
	rm -rf build

test:
	go test $(TEST_FLAGS) ./...

