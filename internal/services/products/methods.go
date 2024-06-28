package products

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"server/internal/domein/models"
)

func (p *Product) CreateUser(ctx context.Context,
	req *models.CreateProductReq) (*models.Product, error) {
	const op = "Prodcut.CreateUser"
	log := p.log.With(
		slog.String("op", op),
		slog.String("Name", req.Name),
	)
	log.Info("Createng Product")
	product, err := p.productSaver.SaveProduct(ctx, &models.CreateProductReq{
		Name:     req.Name,
		Price:    req.Price,
		Category: req.Category,
		Count:    req.Count,
	})
	if err != nil {
		return nil, ErrInvalitArguments
	}
	log.Info("Product succsesfuly saved")
	fmt.Println(product)
	return product, nil
}

func (p *Product) DeleteProduct(ctx context.Context, req *models.ProductDeleteReq) (*models.Product, error) {
	const op = "Product.DeleteProduct"
	log := p.log.With(
		slog.String("op", op),
		slog.String("ID", req.ID),
	)
	log.Info("Deleting Product")
	product, err := p.productDeleter.DeleteProduct(ctx, &models.ProductDeleteReq{
		ID:           req.ID,
		IsHardDelete: req.IsHardDelete,
	})
	if err != nil {
		if errors.Is(err, ErrProductNotFound) {
			return nil, err
		}
	}
	return product, nil
}

func (p *Product) GetAllProduct(ctx context.Context,
	req *models.GetAllProductReq) ([]*models.Product, error) {
	const op = "Product.GetAllProduct"
	log := p.log.With(
		slog.String("op", op),
		slog.String("Req.Field", req.Field),
	)
	log.Info("GetAllProduct Product")
	products, err := p.prodcutProvider.GetAllProduct(ctx, req)
	if err != nil {
		if errors.Is(err, ErrInvalitArguments) {
			return nil, ErrInvalitArguments
		}
		return nil, err
	}
	log.Info("GetAllProduct Succesfully")
	return products, nil
}

func (p *Product) GetProductById(ctx context.Context, id string) (*models.Product, error) {
	const op = "Product.GetProductById"
	log := p.log.With(
		slog.String("op", op),
		slog.String("Req.ID", id),
	)
	log.Info("Get product is progress")
	product,err := p.prodcutProvider.GetProdcutById(ctx,id)
	if err != nil {
		if errors.Is(err,ErrProductNotFound){
			return nil,ErrProductNotFound
		}

		return nil,err
	}
	log.Info("Product succesfuly getting")

	return product,nil
}


func (p *Product)UpdateProduc(ctx context.Context, 
	req *models.UpdateProducReq)(*models.Product,error) {
		const op = "Product.UpdateProduc"
		log := p.log.With(
			slog.String("op", op),
			slog.String("Req",fmt.Sprintf("%s %s %s %d %s",req.Name,
			req.Category,req.Price,req.Count,req.ID)),
		)
		log.Info("Upateng is proggress")
		product,err := p.productUpdater.UpdateProduct(ctx,req)
		if err != nil {
			if errors.Is(err,ErrProductNotFound) {
				return nil,ErrProductNotFound
			}
			return nil,err
		}
		log.Info("Succesfuly updated ")

		return product,nil
}