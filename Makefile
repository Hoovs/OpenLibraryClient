TEST_FLAGS := -v -race

vendor:
	dep ensure

build:
	mkdir build

build/OpenLibrary: build
	go build -o build/OpenLibraryServer ./server

cleanBuild:
	rm -rf build

test:
	go test $(TEST_FLAGS) ./...

wishList.sqlite3:
	sqlite3 wishList.sqlite3 < sql/wishlist.sql

image: build/OpenLibrary wishList.sqlite3
