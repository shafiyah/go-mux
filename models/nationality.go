package models

type Nationality struct {
	NationalityID   uint   `gorm:"primaryKey" json:"id"`
	NationalityName string `gorm:"type:varchar(50); not null" json:"kewarganegaraan"`
	NationalityCode string `gorm:"type:char(2); not null" json:"kode"`
}
