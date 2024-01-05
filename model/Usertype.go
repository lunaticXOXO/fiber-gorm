package model


type Usertype struct {

	IdType int 		   `gorm:"type:int(11);primaryKey" json:"idtype"`
	Description string `gorm:"type:varchar(255)" json:"description"`
}