package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignKey:Cid"`
	Title    string   `gorm:"column:title;type:varchar(100);not null" json:"title"`
	Cid      int      `gorm:"column:cid;type:int;not null" json:"cid"`
	Desc     string   `gorm:"column:desc;type:varchar(200)" json:"desc"`
	Content  string   `gorm:"column:content;type:longtext" json:"content"`
	Img      string   `gorm:"column:img;type:varchar(100)" json:"img"`
}

func (Article) TableName() string {
	return "article"
}
