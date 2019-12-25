package Auth

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"iris/config"
	"time"
)

func TokenHandler(ctx iris.Context) string {
	return GetToken(1, 1)
}

func GetToken(Id int64, RoleId int64) string {
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": Id,
		"role_id": RoleId,
		"iss":     "iris",                                                  // 签发人
		"iat":     time.Now().Unix(),                                       // 签发时间
		"exp":     time.Now().Add(1 * time.Hour * time.Duration(1)).Unix(), // 设定过期时间，便于测试，设置1分钟过期
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString([]byte(config.GetJwt().Secret))
	return tokenString
}
