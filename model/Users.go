package model

type Users struct{
	
	Username string 	`gorm:"type:varchar(15)" json:"username"`
	Password string 	`gorm:"type:varchar(15)" json:"password"`
	Type     int 		`gorm:"type:int(11)" json:"type"`
	Usertype Usertype 	`gorm:"foreignKey:Type;references:IdType"`
}