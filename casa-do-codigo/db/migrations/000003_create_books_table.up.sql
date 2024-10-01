BEGIN;

CREATE TABLE IF NOT EXISTS books (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  abstract TEXT NOT NULL,
  table_of_content TEXT NOT NULL,
  price FLOAT NOT NULL,
  number_of_pages INT NOT NULL,
  isbn VARCHAR(255) NOT NULL,
  publish_date DATE NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  category_id INT,
  author_id INT NOT NULL,

  FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE SET NULL,
  FOREIGN KEY (author_id) REFERENCES authors (id) ON DELETE CASCADE
);

COMMIT;