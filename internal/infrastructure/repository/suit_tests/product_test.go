package suittests

import (
	"context"
	"fmt"
	"log"
	"server/internal/domein/models"
	repo "server/internal/infrastructure/repository/postgresql"
	configpkg "server/internal/pkg/config"
	"server/internal/pkg/postgres"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ProductTest struct {
	suite.Suite
	CleanUpFunc func()
	Repository  *repo.ProductRepository
}

func (s *ProductTest) SetupTest() {
	cfg, err := configpkg.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	pgPool, err := postgres.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	s.Repository = repo.NewProductRepository(pgPool)

	s.CleanUpFunc = func() {

	}
}

func (s *ProductTest) TearDownTest() {
	if s.CleanUpFunc != nil {
		s.CleanUpFunc()
	}
}
func TestProductTestSuite(t *testing.T) {
	suite.Run(t, new(ProductTest))
}
func (p *ProductTest) TestProduct() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	productCreateRequest := models.CreateProductReq{
		Name:     "shaptoli",
		Price:    "1234$",
		Category: "MEVALAR",
		Count:    12,
	}

	productCreateRes, err := p.Repository.SaveProduct(ctx, &productCreateRequest)
	p.Require().NotNil(productCreateRes)
	p.Require().NoError(err)
	p.Equal(productCreateRequest.Name, productCreateRes.Name)
	p.Equal(productCreateRequest.Category, productCreateRes.Category)
	p.Equal(productCreateRequest.Count, productCreateRes.Count)
	p.Equal(productCreateRequest.Price, productCreateRes.Price)

	getResponse, err := p.Repository.GetProdcutById(ctx, productCreateRes.Id)
	p.NotNil(getResponse)
	p.Require().NoError(err)
	p.Equal(productCreateRequest.Name, getResponse.Name)
	p.Equal(productCreateRequest.Category, getResponse.Category)
	p.Equal(productCreateRequest.Count, getResponse.Count)
	p.Equal(productCreateRequest.Price, getResponse.Price)

	productsAll, err := p.Repository.GetAllProduct(ctx, &models.GetAllProductReq{
		Field: "name_product",
		Value: "shaptoli",
	})
	p.NotNil(productsAll)
	p.NoError(err)
	updateProduct, err := p.Repository.UpdateProduct(ctx, &models.UpdateProducReq{
		ID:       productCreateRes.Id,
		Name:     "ANOR",
		Price:    "tEST",
		Category: "MEVALAR",
		Count:    12,
	})
	p.NoError(err)
	p.NotNil(updateProduct)
	fmt.Println(updateProduct)

	deleteProduct, err := p.Repository.DeleteProduct(ctx, &models.ProductDeleteReq{
		IsHardDelete: true,
		ID:           productCreateRes.Id,
	})
	p.NotNil(deleteProduct)
	p.NoError(err)

}
