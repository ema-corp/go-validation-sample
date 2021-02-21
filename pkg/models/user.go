package models

import _ "github.com/go-playground/validator"

type User struct {
    Name   string  `validate:"required"` // :（セミコロン）の前後にスペースを開けない
    IconID int  `validate:"is_icon_id,required"`
}

type UserJson struct {
    Name string `json:"name" validate:"required"` // tag を共存させるときはスペースを使う
}

