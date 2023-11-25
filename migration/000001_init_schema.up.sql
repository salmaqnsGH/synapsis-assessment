CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    username VARCHAR(255),
    password VARCHAR(255),
    balance bigint,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description VARCHAR(255),
    createdAt TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deletedAt TIMESTAMPTZ
);

CREATE TABLE products (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255),
    description VARCHAR(255),
	category_id INT REFERENCES categories(id),
	owner_id INT REFERENCES users(id),
    quantity INT,
    price BIGINT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);


INSERT INTO categories (name, description) VALUES
('Electronics', 'Description for Category 1'),
('Beauty', 'Description for Category 2'),
('Fashion', 'Description for Category 3');

