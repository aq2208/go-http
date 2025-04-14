package model

import "time"

type Order struct {
	OrderId int
	UserId int
	OrderDate time.Time
	TotalAmount float32
}