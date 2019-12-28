package Models

import (
	"iris/app/Entry"
	"iris/config"
)

//查
func GetUserInfo(id int64) (user *Entry.User) {
	user = &Entry.User{}
	err := config.DB.First(user, "id = ?", id).Error
	if err != nil {
		config.Log.LogInfo(err)
		return nil
	}
	return user
}

//用户登录获取信息
func GetUserInfoByUserName(name string) (user *Entry.User, err error) {
	user = &Entry.User{}
	err = config.DB.First(user, "username = ?", name).Error
	if err != nil {
		return nil, err
	}
	return user, err
}
