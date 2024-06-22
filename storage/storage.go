package storage

import "context"

type Storage interface {
	CreateProduct(ctx context.Context)
}
