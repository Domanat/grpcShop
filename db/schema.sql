CREATE TABLE products (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price BIGINT NOT NULL CHECK (price >= 0)
);