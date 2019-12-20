package Models

import (
	"iris/app/Entry"
	"iris/config"
)

//查
func GetUserInfo(id int64) (user *Entry.User, err error) {
	user = &Entry.User{}
	err = config.DB.First(user, "id = ?", id).Error
	return user, err
}
