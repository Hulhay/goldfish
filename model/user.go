package model

const (
	SUPER_USER = "super_user"
	PETUGAS    = "petugas"
	WARGA      = "warga"
)

var (
	IsValidRole = map[string]bool{
		SUPER_USER: true,
		PETUGAS:    true,
		WARGA:      true,
	}
)

type User struct {
	ID           int64  `json:"id" gorm:"column:user_id;type:int;primary key;auto_increment"`
	UserName     string `json:"user_name" gorm:"column:user_name;type:varchar(255)"`
	UserUsername string `json:"user_username" gorm:"column:user_username;type:varchar(255)"`
	UserEmail    string `json:"user_email" gorm:"column:user_email;type:varchar(255)"`
	UserRole     string `json:"user_role" gorm:"column:user_role;type:varchar(255)"`
	UserIsLogin  bool   `json:"user_is_login" gorm:"column:user_is_login;"`
	UserPassword string `json:"-" gorm:"column:user_password;"`
}
