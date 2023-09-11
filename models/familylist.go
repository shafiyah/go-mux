package models

type FamilyList struct {
	FlID       uint   `gorm:"primaryKey" json:"id"`
	FlName     string `gorm:"type:varchar(50); not null" json:"nama"`
	FlRelation string `gorm:"type:varchar(50); not null" json:"relasi"`
	FlDob      string `gorm:"type:date; not null" json:"tanggal_lahir"`
	CustomerID uint   `gorm:"not null"`
}
