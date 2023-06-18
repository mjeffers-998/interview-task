package storage

import (
	"time"

	"github.com/mjeffers-998/interview-task/structs"
)

type Storage interface {
	ListAllOrders(status string) ([]*structs.Order, error)
	GetOrderByID(id int64) (*structs.Order, error)
	GetRevenue(startDate time.Time, endTime time.Time) (float32, error)
	CreateOrder(*structs.Order) (int64, error)
	CancelOrder(id int64) error
}
