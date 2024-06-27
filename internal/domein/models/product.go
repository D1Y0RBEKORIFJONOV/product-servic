package models

type Product struct {
	Id string 
	Name string
	Category string
	Price string
	Count int64
	Status string
	Created_at string
	Updated_at string
	Deleted_at string
}


type CreateProductReq struct {
	Name string
	Category string
	Price string
	Count int64
}
type ProductDeleteReq struct {
	ID string
	IsHardDelete bool
}
type Empty struct {}

type GetAllProductReq struct {
	Field string
	Value string
	Limit int32
	Page int32
}

type UpdateProducReq  struct{
	Name string
	Category string
	Price string
	Count int64
}