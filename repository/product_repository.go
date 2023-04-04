package repository

import "latihan/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
}
