CREATE TABLE wishlist (
  id INT PRIMARY KEY,
  userId INT NOT NULL,
  bookId INT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  deleted boolean DEFAULT false
);