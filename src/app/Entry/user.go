package Entry

import "time"

type User struct {
	Id        int64     `gorm:"column:id;primary_key" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	UserName  string    `gorm:"column:username" json:"username"`
	Password  string    `gorm:"column:password" json:"password"`
	RoleId    int64     `gorm:"column:role_id" json:"role_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

func (e *User) TableName() string {
	return "users"
}
