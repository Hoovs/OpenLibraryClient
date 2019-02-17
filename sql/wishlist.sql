CREATE TABLE wishlist (
  id INTEGER PRIMARY KEY,
  userId INT NOT NULL,
  bookTitle TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  deleted boolean DEFAULT false
);