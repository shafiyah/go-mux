package models

type Customer struct {
	CustomerID    uint   `gorm:"primaryKey" json:"id"`
	CstName       string `gorm:"type:varchar(50); not null" json:"name"`
	CstEmail      string `gorm:"type:varchar(50); not null" json:"email"`
	CstPhoneNum   string `gorm:"type:varchar(20); not null" json:"phone"`
	CstDob        string `gorm:"type:date; not null" json:"tanggal_lahir"`
	NationalityID uint   `gorm:"not null"`

	Nationality *Nationality
	FamilyLists []FamilyList `gorm:"foreignKey:CustomerID"`
}
