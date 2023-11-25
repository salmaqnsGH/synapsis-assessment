CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    username VARCHAR(255),
    password VARCHAR(255)
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description VARCHAR(255)
);

INSERT INTO categories (name, description) VALUES
('Electronics', 'Description for Category 1'),
('Beauty', 'Description for Category 2'),
('Fashion', 'Description for Category 3');
