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
    quantity INT,
    price BIGINT,
	category_id INT REFERENCES categories(id),
	owner_id INT REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    quantity INT,
    price BIGINT,
    total_price BIGINT, 
    is_in_cart BOOLEAN DEFAULT TRUE,
    user_id INT REFERENCES users(id),
    product_id INT REFERENCES products(id),
    owner_id INT REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ 
);



INSERT INTO categories (name, description) VALUES
('Electronics', 'Gadgets, devices, and other electronic equipment'),
('Books', 'All kinds of books and literature'),
('Clothing', 'Apparel for men, women, and children');
