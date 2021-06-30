package model

import (
	"gorm.io/gorm"
	"mofan-blog/utils/errmsg"
)

type Category struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(20);not null" json:"name"`
}

func (Category) TableName() string {
	return "category"
}

// CheckCategory 查询分类是否存在
func CheckCategory(name string) int {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

// CreateCate 新增分类
func CreateCate(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCateInfo 查询单个分类信息
func GetCateInfo(id int) (Category, int) {
	var cate Category
	db.Where("id = ?", id).First(&cate)
	return cate, errmsg.SUCCESS
}

// GetCate 查询分类列表
func GetCate(pageSize, pageNum int) ([]Category, int64) {
	var categories []Category
	var total int64
	err := db.Model(&Category{}).Find(&categories).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return categories, total
}

// EditCate 编辑用户信息
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := db.Model(cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteCate 删除用户
func DeleteCate(id int) int {
	err := db.Where("id = ?", id).Delete(&Category{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
