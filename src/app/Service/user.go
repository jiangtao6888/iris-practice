package Service

import (
	"iris/app/Entry"
	"iris/app/Models"
)

var User = &user{}

type user struct{}

func (s *user) GetUserInfo(userId int64) (*Entry.User, error) {
	user, err := Models.GetUserInfo(userId)
	return user, err
}
