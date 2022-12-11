package model

const (
	ADMIN   = "admin"
	PETUGAS = "petugas"
	WARGA   = "warga"
)

var (
	IsValidRole = map[string]bool{
		ADMIN:   true,
		PETUGAS: true,
		WARGA:   true,
	}
)

type User struct {
	ID       int64  `json:"id" gorm:"column:user_id;type:int;primary key;auto_increment"`
	Name     string `json:"name" gorm:"column:user_name;type:varchar(255)"`
	Email    string `json:"email" gorm:"column:user_email;type:varchar(255)"`
	Role     string `json:"role" gorm:"column:user_role;type:varchar(255)"`
	IsLogin  bool   `json:"is_login" gorm:"column:user_is_login;"`
	Password string `json:"-" gorm:"column:user_password;"`
}
