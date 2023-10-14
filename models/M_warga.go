package models


type Wargas struct {

    ID            uint   `gorm:"primaryKey" json:"id" form:"id"`
    Nik           string `gorm:"not null;unique" json:"nik"`
    Nama          string `gorm:"not null" json:"nama"`
    Umur          uint `gorm: json:"umur"`
    Id_Pekerjaan   uint `gorm:"foreignKey:ID" json:"id_pekerjaan"`
    Alamat        string `gorm:"not null;" json:"Alamat"`

}
