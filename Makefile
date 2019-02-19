TEST_FLAGS := -v

vendor:
	dep ensure

build:
	mkdir build

build/OpenLibrary: build
	env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags "-linkmode external -extldflags -static" -o build/OpenLibraryServer ./server

cleanBuild:
	rm -rf build

test:
	go test $(TEST_FLAGS) ./...

wishList.sqlite3:
	sqlite3 wishList.sqlite3 < sql/wishlist.sql

image: build/OpenLibrary wishList.sqlite3
	docker build -t open_library:latest .

imageNoGo:
	docker build -t open_library:latest -f Dockerfile_complete .

up:
	docker-compose up

down:
	docker-compose down
