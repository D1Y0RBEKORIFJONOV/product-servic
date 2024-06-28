package suittests

import (
	"log"
	repo "server/internal/infrastructure/repository/postgresql"
	configpkg "server/internal/pkg/config"
	"server/internal/pkg/postgres"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ProductTest struct {
	suite.Suite
	CleanUpFunc func()
	Repository  repo.ProductRepository
}

func (s *ProductTest) SetupTest() {
	cfg, err := configpkg.NewConfig()
	if err != nil {
		log.Println(err)
		return
	}
	pgPool, err := postgres.New(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	s.Repository = *repo.NewProductRepository(pgPool)

}
func (s *ProductTest) TearDownTest() {
	s.CleanUpFunc()
}

func TestProductTestSuite(t *testing.T) {
	suite.Run(t, new(ProductTest))
}
