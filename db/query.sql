-- name: GetProduct :one
SELECT id, name, price
FROM products
WHERE id = $1;

-- name: CreateProduct :one
INSERT INTO products (name, price) 
VALUES ($1, $2) RETURNING id;

-- name: ListProducts :many
SELECT id, name, price
FROM products 
ORDER BY id;

-- name: DeleteProduct :exec
DELETE FROM products 
WHERE id = $1;