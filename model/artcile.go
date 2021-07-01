package model

import (
	"gorm.io/gorm"
	"mofan-blog/utils/errmsg"
)

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

// CreateArt 新增文章
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCateArt 查询分类下的所有文章
func GetCateArt(id, pageSize, pageNum int) ([]Article, int, int64) {
	var cateArtList []Article
	var total int64
	// todo 软删除的category存在一定的问题
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&cateArtList).Count(&total).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCESS, total
}

// GetArtInfo 查询单个文章信息
func GetArtInfo(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}

// GetArt 查询文章列表
func GetArt(title string, pageSize, pageNum int) ([]Article, int, int64) {
	var articles []Article
	var total int64
	if title == "" {
		err := db.Model(&Article{}).Preload("Category").Order("Updated_At DESC").Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, errmsg.ERROR, 0
		}
		return articles, errmsg.SUCCESS, total
	}
	err := db.Model(&Article{}).Preload("Category").Order("Updated_At DESC").Model(&Article{}).Where("title LIKE ?", title+"%").Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articles, errmsg.SUCCESS, total
}

// EditArt 编辑用户信息
func EditArt(id int, data *Article) int {
	var cate Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err := db.Model(cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteArt 删除文章
func DeleteArt(id int) int {
	err := db.Where("id = ?", id).Delete(&Article{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
