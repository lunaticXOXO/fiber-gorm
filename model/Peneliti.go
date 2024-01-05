package model

type Peneliti struct {
	Nidn            string         	`gorm:"primaryKey;type:varchar(255)" json:"nidn"`
	Nama            string         	`gorm:"type:varchar(255)" json:"nama"`
	Email           string         	`gorm:"type:varchar(255)" json:"email"`
	Telephone       string         	`gorm:"type:varchar(255)" json:"telephone"`
	Anggota         string         	`gorm:"type:varchar(255)" json:"anggota"`
	RisetID         string         	`gorm:"type:varchar(255)" json:"risetid"`
	RisetPenelitian RisetPenelitian `gorm:"foreignKey:RisetID;references:Idriset"`

}