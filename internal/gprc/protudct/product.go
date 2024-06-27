package protudct

import (
	"server/internal/domein/models"
	"server/protductDB"


	"google.golang.org/grpc"
)

type serverApi struct {
	protductDB.UnimplementedProductServerServer
	product       Product
	statusProduct map[string]models.Product
}

func Register(GRPC *grpc.Server, product Product) {
	protductDB.RegisterProductServerServer(GRPC,
		&serverApi{
			statusProduct: make(map[string]models.Product),
			product:       product,
		},
	)
}
