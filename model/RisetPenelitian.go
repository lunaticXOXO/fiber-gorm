package model

type RisetPenelitian struct {

	Idriset string `gorm:"primarykey;type:varchar(255)" json:"idriset"`
	Judul string `gorm:"type:varchar(255)" json:"judul"`
	No_Kontrak string `gorm:"type:varchar(255)" json:"no_kontrak"`
	Tahun string `gorm:"type:varchar(255)" json:"tahun"`
	Anggaran int `gorm:"type:int(11)" json:"anggaran"`

}