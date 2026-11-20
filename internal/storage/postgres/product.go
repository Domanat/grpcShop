package postgres

import (
	"context"
	"fmt"
	"log"

	"shop/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	pool *pgxpool.Pool
}

func NewProductRepository(pool *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{pool: pool}
}

func (p *ProductRepository) Create(ctx context.Context, name string, price int64) (int64, error) {
	query := `INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id`

	var newId int64

	err := p.pool.QueryRow(ctx, query, name, price).Scan(newId)
	if err != nil {
		return -1, fmt.Errorf("Error inserting into database %w\n", err)
	}

	return newId, nil
}

func (p *ProductRepository) GetById(ctx context.Context, id int64) (model.Product, error) {
	var product model.Product

	err := p.pool.QueryRow(
		ctx,
		`SELECT * from products WHERE id = $1`,
		id,
	).Scan(product)

	log.Println(product)

	return product, err
}

func (p *ProductRepository) GetList(ctx context.Context, limit int) ([]model.Product, error) {
	query := `SELECT id, name, price FROM products ORDER BY id ASC LIMIT $1`

	rows, err := p.pool.Query(ctx, query, limit)
	if err != nil {
		return nil, fmt.Errorf("Query failed: %w\n", err)
	}

	defer rows.Close()
	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Product])
	if err != nil {
		return nil, fmt.Errorf("Error collecting rows %w", err)
	}

	return products, nil
}
