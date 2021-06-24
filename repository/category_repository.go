package repository

import "github.com/bryanagamk/golang-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
