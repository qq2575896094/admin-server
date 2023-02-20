package model

type Meta struct {
	CreateAt int64
	UpdateAt int64
}

type User struct {
	Id          string `bson:"_id"`
	Username    string
	Password    string
	Avatar      string
	Type        string
	PhoneNumber string
	RealName    string
	Gender      string
	Age         int
	Meta
}
