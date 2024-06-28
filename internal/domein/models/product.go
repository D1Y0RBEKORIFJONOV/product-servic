package models

import "time"

type Product struct {
	Id         string
	Name       string
	Category   string
	Price      string
	Count      int64
	Status     string
	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}

type CreateProductReq struct {
	Name     string
	Category string
	Price    string
	Count    int64
}
type ProductDeleteReq struct {
	ID           string
	IsHardDelete bool
}
type Empty struct{}

type GetAllProductReq struct {
	Field string
	Value string
	Limit uint64
	Page  uint64
}

type UpdateProducReq struct {
	ID       string
	Name     string
	Category string
	Price    string
	Count    int64
}
