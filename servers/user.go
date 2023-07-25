package servers

import (
	"context"
	"errors"
	"github.com/qq2575896094/admin-server/constants"
	"github.com/qq2575896094/admin-server/dao"
	"github.com/qq2575896094/admin-server/models"
	"github.com/qq2575896094/admin-server/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserServer struct {
	ctx context.Context
}

// SignUp 注册
func (u *UserServer) SignUp(user *models.UserRegisterParams) (*models.UserInfo, error) {
	now := time.Now()
	user.CreateAt = now
	user.UpdateAt = now
	user.Avatar = constants.UserAvatarDefault
	user.Gender = constants.UserGenderDefault
	user.Type = constants.UserTypeDefault

	password, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = password

	result, err := dao.AddUser(u.ctx, user)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("user with that already exist")
		}
		return nil, err
	}

	var userInfo models.UserInfo
	err = dao.GetUserById(u.ctx, result.InsertedID, &userInfo)
	if err != nil {
		return nil, err
	}

	return &userInfo, nil
}

// Login 登录
func (u *UserServer) Login(user *models.UserLoginParams) (*models.UserInfo, error) {
	var userInfo models.UserInfo
	err := dao.GetUserByName(u.ctx, user.Username, &userInfo)
	if err != nil {
		return nil, errors.New("user not fond")
	}

	err = utils.ComparePassword(userInfo.Password, user.Password)
	if err != nil {
		return nil, errors.New("password is not correct")
	}

	return &userInfo, nil
}

func NewUserServer() *UserServer {
	return &UserServer{}
}
