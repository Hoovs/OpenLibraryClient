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
 - [ ] Wish list allows Addition of book.
 - [ ] Wish list allows Delete of a book.
 - [ ] Unit tests are completed.
 - [ ] Document database choice reason.
 - [ ] (Optional) Add ability for read list.
 
## Prerequisites
 - Must have sqlite3 installed locally.