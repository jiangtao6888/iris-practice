package Service

import (
	"encoding/json"
	"iris/app/Code"
	"iris/app/Entry"
	"iris/app/Exceptions"
	"iris/app/Models"
	"iris/config"
	"time"
)

var User = &user{}

type user struct{}

func (s *user) GetUserInfo(UserId int64) (user *Entry.User, err error) {

	pre := "user_info_key"
	key := config.Cache.GetKey(pre, UserId)
	val := config.Cache.GetArray(key)

	if val == "" {
		user = Models.GetUserInfo(UserId)
		if user != nil {
			err := config.Cache.SetArray(key, user, 300*time.Second)
			if err != nil {
				return nil, Exceptions.New(Code.ErrorCode, "set redis fail")
			}
		} else {
			return nil, Exceptions.New(Code.ErrorCode, "set redis fail")
		}
	}
	err = json.Unmarshal([]byte(val), user)
	if err != nil {
		return nil, err
	}
	return user, nil

}
