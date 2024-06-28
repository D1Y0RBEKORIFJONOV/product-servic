package products

import (
	"errors"
	"log/slog"
)

type Product struct {
	log             *slog.Logger
	productSaver    SaverProduct
	productDeleter  DeleterProduct
	prodcutProvider ProviderProduct
	productUpdater  UpdaterProduct
}

func NewProduct(
	log *slog.Logger,
	productSaver SaverProduct,
	productDeleter DeleterProduct,
	productProvider ProviderProduct,
	productUpdater UpdaterProduct,
) *Product {
	return &Product{
		log:             log,
		productSaver:    productSaver,
		productDeleter:  productDeleter,
		productUpdater:  productUpdater,
		prodcutProvider: productProvider,
	}
}

var (
	ErrInvalitArguments = errors.New("invalid arguments")
	ErrProductNotFound  = errors.New("product not found")
)

