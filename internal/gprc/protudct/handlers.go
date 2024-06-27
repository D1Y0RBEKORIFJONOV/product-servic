package protudct

import (
	"context"
	"server/internal/domein/models"
	"server/protductDB"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverApi) CreateProduct(ctx context.Context,
	req *protductDB.CreateProductReq) (*protductDB.Product, error) {
	if req.Category == "" || req.Count == 0 || req.Name == "" || req.Price == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid arguments")
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second*4)
	defer cancel()
	product, err := s.product.CreateUser(ctx, &models.CreateProductReq{
		Name:     req.Name,
		Price:    req.Price,
		Category: req.Category,
		Count:    req.Count,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	s.statusProduct[product.Id] = *product
	return &protductDB.Product{
		Id:        product.Id,
		Name:      product.Name,
		Price:     product.Price,
		Count:     product.Count,
		Status:    product.Status,
		Category:  product.Category,
		CreatedAt: product.Created_at,
		UpdatedAt: product.Updated_at,
		DeletedAt: product.Deleted_at,
	}, nil
}

func (s *serverApi) DeletedProduct(ctx context.Context, req *protductDB.DeleteReq) (*protductDB.Empty, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid arguments")
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second*4)
	defer cancel()
	product, err := s.product.DeleteProduct(ctx, &models.ProductDeleteReq{
		ID:           req.Id,
		IsHardDelete: req.IsHardDelete,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	s.statusProduct[product.Id] = *product

	return nil, nil
}

func (s *serverApi) GetAllProducts(ctx context.Context, req *protductDB.GetAllProductReq) (*protductDB.GetAllProductsRes, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*4)
	defer cancel()
	products, err := s.product.GetAllProduct(ctx, &models.GetAllProductReq{
		Field: req.Field,
		Value: req.Value,
		Limit: req.Limit,
		Page:  req.Page,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var resPorduct []*protductDB.Product

	for i := 0; i < len(products); i++ {
		var product protductDB.Product
		product.Id = products[i].Id
		product.Name = products[i].Name
		product.Category = products[i].Category
		product.Price = products[i].Price
		product.Count = products[i].Count
		product.Status = products[i].Status
		product.CreatedAt = products[i].Created_at
		product.UpdatedAt = products[i].Updated_at
		product.DeletedAt = products[i].Deleted_at

		resPorduct = append(resPorduct, &product)
	}
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	return &protductDB.GetAllProductsRes{
		Products: resPorduct,
	}, nil
}

func (s *serverApi) GetProductById(ctx context.Context, req *protductDB.ProductByIdReq) (*protductDB.Product, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid arguments")
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second*4)
	defer cancel()
	product, err := s.product.GetProductById(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	return &protductDB.Product{
		Id:        product.Id,
		Name:      product.Name,
		Price:     product.Price,
		Count:     product.Count,
		Status:    product.Status,
		Category:  product.Category,
		CreatedAt: product.Created_at,
		UpdatedAt: product.Updated_at,
		DeletedAt: product.Deleted_at,
	}, nil
}

func (s *serverApi) ShowRealTimeAddinAndDeleteing(_ *protductDB.Empty,
	stream protductDB.ProductServer_ShowRealTimeAddinAndDeleteingServer) error {
	for {
		if len(s.statusProduct) > 0 {
			for key, product := range s.statusProduct {
				stream.SendMsg(product.Status)
				stream.Send(&protductDB.Product{
					Id:        product.Id,
					Name:      product.Name,
					Price:     product.Price,
					Count:     product.Count,
					Status:    product.Status,
					Category:  product.Category,
					CreatedAt: product.Created_at,
					UpdatedAt: product.Updated_at,
					DeletedAt: product.Deleted_at,
				})
				delete(s.statusProduct,key)
			}
		}
	}
}

func (s *serverApi) UpdateProduc(ctx context.Context, req *protductDB.UpdateProductReq) (*protductDB.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*4)
	defer cancel()
	product,err := s.product.UpdateProduc(ctx,&models.UpdateProducReq{
		Name: req.Name,
		Category: req.Category,
		Count: req.Count,
		Price: req.Proce,
	})
	if err!= nil {
		return nil,status.Error(codes.Internal,err.Error())
	}
	return &protductDB.Product{
		Id:        product.Id,
		Name:      product.Name,
		Price:     product.Price,
		Count:     product.Count,
		Status:    product.Status,
		Category:  product.Category,
		CreatedAt: product.Created_at,
		UpdatedAt: product.Updated_at,
		DeletedAt: product.Deleted_at,
	},nil
}
