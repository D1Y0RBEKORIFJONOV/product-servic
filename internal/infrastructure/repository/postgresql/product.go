package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"server/internal/domein/models"
	"server/internal/pkg/postgres"
	"server/internal/services/products"
	"time"
)

// type UpdaterProduct interface {
// 	UpdateProduct(ctx context.Context,
// 		req *models.UpdateProducReq) (*models.Product,error)
// }

type ProductRepository struct {
	db        *postgres.PostgresDB
	tableName string
}

func NewProductRepository(pg *postgres.PostgresDB) *ProductRepository {
	return &ProductRepository{
		db:        pg,
		tableName: "products",
	}
}

func (p *ProductRepository) selectQueryProduct() string {
	return `
	id,
	name_product,
	category,
	price,
	count,
	status_product,
	created_at,
	updated_at,
	deleted_at
	`
}
func (p *ProductRepository) SaveProduct(ctx context.Context, req *models.CreateProductReq) (*models.Product, error) {
	data := map[string]interface{}{
		"name_product":   req.Name,
		"category":       req.Category,
		"price":          req.Price,
		"count":          req.Count,
		"status_product": "created",
	}
	query, args, err := p.db.Sq.Builder.Insert(p.tableName).
		SetMap(data).
		Suffix(fmt.Sprintf("RETURNING %s", p.selectQueryProduct())).ToSql()
	if err != nil {
		return nil, err
	}

	var product models.Product
	var updatedAt, deletedAt sql.NullTime

	err = p.db.QueryRow(ctx, query, args...).Scan(
		&product.Id,
		&product.Name,
		&product.Category,
		&product.Price,
		&product.Count,
		&product.Status,
		&product.Created_at,
		&updatedAt,
		&deletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductRepository) GetProdcutById(ctx context.Context, id string) (*models.Product, error) {
	query, argc, err := p.db.Sq.Builder.
		Select(p.selectQueryProduct()).
		From(p.tableName).Where(p.db.Sq.Equal("id", id)).ToSql()

	if err != nil {
		return nil, err
	}

	var product models.Product
	var updatedAt, deletedAt sql.NullTime
	err = p.db.QueryRow(ctx, query, argc...).Scan(
		&product.Id,
		&product.Name,
		&product.Category,
		&product.Price,
		&product.Count,
		&product.Status,
		&product.Created_at,
		&updatedAt,
		&deletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, products.ErrProductNotFound
		}
		return nil, err
	}

	return &product, nil
}

func (p *ProductRepository) DeleteProduct(ctx context.Context, req *models.ProductDeleteReq) (*models.Product, error) {
	product, err := p.GetProdcutById(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	var (
		query string
		args  []interface{}
	)
	if req.IsHardDelete {
		query, args, err = p.db.Sq.Builder.Delete(p.tableName).Where(p.db.Sq.Equal("id", product.Id)).ToSql()
	} else {
		query, args, err = p.db.Sq.Builder.Update(p.tableName).Set("updated_at", time.Now()).
			Where(p.db.Sq.Equal("id", product.Id)).ToSql()
	}
	if err != nil {
		return nil, err
	}
	r, err := p.db.Exec(ctx, query, args...)

	if err != nil {
		return nil, err
	}
	if r.RowsAffected() == 0 {
		return nil, products.ErrProductNotFound
	}
	return product, nil
}

func (p *ProductRepository) GetAllProduct(ctx context.Context,
	req *models.GetAllProductReq) ([]*models.Product, error) {

	toSql := p.db.Sq.Builder.Select(p.selectQueryProduct()).From(p.tableName)
	if req.Field != "" && req.Value != "" {
		toSql = toSql.Where(p.db.Sq.Equal(req.Field, req.Value))
	}
	if req.Page != 0 {
		toSql = toSql.Offset(req.Page)
	}
	if req.Limit != 0 {
		toSql = toSql.Limit(req.Limit)
	}

	query, args, err := toSql.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []*models.Product
	for rows.Next() {
		var product models.Product
		var updatedAt, deletedAt sql.NullTime
		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Category,
			&product.Price,
			&product.Count,
			&product.Status,
			&product.Created_at,
			&updatedAt,
			&deletedAt)

		if err != nil {
			return nil, err

		}
		if deletedAt.Valid {
			product.Deleted_at = deletedAt.Time
		}
		if updatedAt.Valid {
			product.Updated_at = updatedAt.Time
		}
		if product.Deleted_at.IsZero() {
			products = append(products, &product)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepository) updatedQuery(req *models.UpdateProducReq) map[string]interface{} {
	return map[string]interface{}{
		"id":             req.ID,
		"name_product":   req.Name,
		"category":       req.Category,
		"price":          req.Price,
		"count":          req.Count,
		"status_product": "updated",
		"updated_at":     time.Now(),
	}
}

func (p *ProductRepository) UpdateProduct(ctx context.Context,
	req *models.UpdateProducReq) (*models.Product, error) {
	data := p.updatedQuery(req)
	query, argc, err := p.db.Sq.Builder.Update(p.tableName).SetMap(data).
		Where(p.db.Sq.Equal("id", req.ID)).
		Suffix(fmt.Sprintf("RETURNING %s", p.selectQueryProduct())).ToSql()
	if err != nil {
		return nil, err
	}
	var deletedAt sql.NullTime
	var product models.Product
	err = p.db.QueryRow(ctx, query, argc...).Scan(
		&product.Id,
		&product.Name,
		&product.Category,
		&product.Price,
		&product.Count,
		&product.Status,
		&product.Created_at,
		&product.Updated_at,
		&deletedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, products.ErrProductNotFound
		}

		return nil, err
	}
	if deletedAt.Valid {
		product.Deleted_at = deletedAt.Time
	}

	return &product, nil
}
