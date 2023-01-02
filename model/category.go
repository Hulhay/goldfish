package model

type Category struct {
	ID            int64  `json:"id" gorm:"column:category_id;type:int;primary key;auto_increment"`
	CategoryName  string `json:"category_name" gorm:"column:category_name;type:varchar(255)"`
	CategoryValue string `json:"category_value" gorm:"column:category_value;type:varchar(255)"`
}
