package model

import (
	"time"
)

type Blog struct {
	Id        string    `json:"id,omitempty" bson:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdateAt  time.Time `json:"updateAt,omitempty"`
}
