-- +migrate Up
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200),
    description VARCHAR(255),
    image_url VARCHAR(255),
    release_year INT,
    price INT,
    total_page INT,
    thickness VARCHAR(10),
    category_id INT REFERENCES categories(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100),
    modified_at TIMESTAMP,
    modified_by VARCHAR(100)
);

-- -- +migrate Down
-- DROP TABLE books;
