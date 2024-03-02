package model

import "time"

type Goods struct {
	ID          int
	ProjectID   int
	Name        string
	Description string
	Priority    int
	Removed     bool
	Created_at  time.Time
}

func NewGoods() *Goods {
	return &Goods{}
}

type DeleteGoods struct {
	ID        int
	ProjectID int
	Removed   int
}

func NewDeleteGoods() *DeleteGoods {
	return &DeleteGoods{}
}

type GoodsSL []*Goods

type Meta struct {
	Total   int
	Removed int
	Limit   int
	Offset  int
}

type List struct {
	Meta
	GoodsSL
}

func NewList() *List {
	return &List{}
}

type Project struct {
	ID         int
	Name       string
	Created_at time.Time
}
