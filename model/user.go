package model

import (
	"time"
)

type User struct {
	Id         string    `json:"id,omitempty" bson:"id,omitempty"`
	Key        string    `json:"key,omitempty" bson:"key,omitempty"`
	Name       string    `json:"name,omitempty" bson:"name,omitempty"`
	Email      string    `json:"email,omitempty" bson:"email,omitempty"`
	OtherEmail string    `json:"otheremail,omitempty" bson:"otheremail,omitempty"`
	Password   string    `json:"password,omitempty" bson:"password,omitempty"`
	LoggedAt   time.Time `json:"loggedat,omitempty" bson:"loggedat,omitempty"`
	CreatedAt  time.Time `json:"createdat,omitempty" bson:"createdat,omitempty"`
	UpdatedAt  time.Time `json:"updatedat,omitempty" bson:"updatedat,omitempty"`
}

type UserViewRes struct {
	Err      int    `json:"err"`
	Emessage string `json:"emessage"`
	Message  string `json:"message"`
	Data     *User  `json:"data"`
	Token    string `json:"token"`
}
