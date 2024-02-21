package usecases

import (
	"cleanr/entities"
)

type OrderRepository interface {
	Save(order entities.Order) error
}
