# Open Library Client
This client extends the [openlibary](https://openlibrary.org/dev/docs/api/search) api 
and adds support for having a wishlist of books to read.

## Design considerations
 - APIs should be RESTful and use the HTTP methods appropriately.
 - You’re using Open Library as a third party client to search and get details of books.
 - You’re building a “wish list” feature as an added feature in the API
 - Design the API endpoints as you see fit but provide clear documentation on how we can
use them.
 - Use any data store you prefer for the wish list functionality (examples can be
Elasticsearch, sqlite, postgres...), just make sure there’s a way for us to use it (be it a
built in DB or an online one)

## Project Features
 - [ ] Fully Documented API interface
 - [X] Wish list allows Addition of book.
 - [X] Wish list allows Delete of a book.
 - [X] Unit tests are completed.
 - [ ] Document database choice reason.
 - [ ] (Optional) Add ability for read list.
 
## Prerequisites
 - Must have sqlite3 installed.
 - Must have docker installed.
 - Must have docker-compose installed.

## How to build
If you want to build the binary instead of using the one specified you need the following:
 - Go 1.10+
 - dep which can be downloaded from `curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh`

You can also use the dockerfile to build this if you rather not install go. 
Instead of following the How to run below run:
```
make imageNoGo
make up
```

## How to run
```$xslt
make vendor
make image
make up
```
Then you can curl as below
```$xslt
curl -XGET http://localhost:8080/wishList/1
curl -XPOST http://localhost:8080/wishList -d '{"userId":1, "bookTitle":"1984"}'
curl -XDELETE http://localhost:8080/wishList/1
curl -XGET http://localhost:8080/search?q=1984
```
