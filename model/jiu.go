package model

import (
	"time"
)

type Jiu struct {
	Id           string `json:"id,omitempty" bson:"id,omitempty"`
	Title        string `json:"title,omitempty"`
	Content      string `json:"content,omitempty"`
	Type         string `json:"type,omitempty"`
	Remark       string `json:"remark,omitempty"`
	Owner        string
	JoinerList   []*Joiner `json:"joinerList,omitempty"`
	PeopleNumber int32     `json:"peoplenumber,omitempty"`
	StartedAt    string    `json:"startedAt,omitempty"`
	EndedAt      string    `json:"endedAt,omitempty"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
	UpdateAt     time.Time `json:"updateAt,omitempty"`
	ImgUrl       []string  `json:"imgurl,omitempty"`
}

type Joiner struct {
	Id   string
	name string
}

type SearchType struct {
	SearchType string `json:"searchType,omitempty"`
}
