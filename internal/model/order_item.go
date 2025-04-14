package model

import "time"

type OrderItem struct {
	OrderId int
	UserId int
	OrderDate time.Time
	TotalAmount float32
}