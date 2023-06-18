package structs

import (
	"fmt"
	"strconv"
	"time"
)

type Order struct {
	OrderID    int `json:"order_id"`
	OrderDate  time.Time
	OrderValue float32
	Status     Status
	StoreID    int
}

func (o *Order) ToStringSlice() []string {
	return []string{strconv.Itoa(o.OrderID), o.OrderDate.GoString(), fmt.Sprintf("%.2f", o.OrderValue), string(o.Status), strconv.Itoa(o.StoreID)}
}

type Status string

const (
	NewOrder       Status = "new"
	CancelledOrder Status = "cancelled"
	Shipped        Status = "shipped"
)
