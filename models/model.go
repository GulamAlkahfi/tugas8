package models


type Pekerjaans struct {
	ID         uint   `gorm:"primaryKey" json:"id" form:"id"`
	Pekerjaan  string `gorm:"not null;unique" json:"pekerjaan"`
}
