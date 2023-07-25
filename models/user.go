package models

import "time"

type Meta struct {
	CreateAt time.Time `json:"createAt" bson:"createAt"`
	UpdateAt time.Time `json:"updateAt" bson:"updateAt"`
}

// UserRegisterParams 用户注册表单
type UserRegisterParams struct {
	Username string `json:"username" bson:"username" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
	Avatar   string `json:"avatar" bson:"avatar"`
	Type     string `json:"type" bson:"type"`
	Gender   string `json:"gender" bson:"gender"`
	Meta
}

// UserLoginParams 用户登录表单
type UserLoginParams struct {
	Username string `json:"username" bson:"username" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

// UserInfo 用户信息
type UserInfo struct {
	Id          string `json:"id" bson:"_id"`
	Username    string `json:"username" bson:"username"`
	RealName    string `json:"realName" bson:"realName"`
	Password    string `json:"-"`
	Avatar      string `json:"avatar" bson:"avatar"`
	Type        string `json:"type" bson:"type"`
	PhoneNumber string `json:"phoneNumber" bson:"phoneNumber"`
	Gender      string `json:"gender" bson:"gender"`
	Age         int    `json:"age" bson:"age"`
	Meta
}
