package model

import (
	"encoding/json"
	"time"
)

type UpdateGoods struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Goods struct {
	ID          int
	ProjectID   int
	Name        string
	Description string
	Priority    int
	Removed     bool
	Created_at  time.Time
}

func (g *Goods) MarshalBinary() ([]byte, error) {
	return json.Marshal(g)
}

type DeleteGoods struct {
	ID        int
	ProjectID int
	Removed   bool
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

type Project struct {
	ID         int
	Name       string
	Created_at time.Time
}

type PriorityGoods struct {
	ID       int
	Priority int `json:"newPriority"`
}

type PriorityGoodsSL []*PriorityGoods

type LogGoods struct {
	ID          int
	ProjectID   int
	Name        string
	Description string
	Priority    int
	Removed     bool
	EventTime   time.Time
}
