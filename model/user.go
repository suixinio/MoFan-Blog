package model

import (
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
	"mofan-blog/utils/errmsg"
)

type User struct {
	gorm.Model
	UserName string `gorm:"column:username;type: varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"column:password;type: varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"column:role;type: int; default:2" json:"role" validate:"required,gte=2"`
}

func (User) TableName() string {
	return "user"
}

//CheckUser 查询用户是否存在
func CheckUser(name string) int {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)

	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// CheckUpUser 查询用户是否存在
func CheckUpUser(id int, name string) int {
	var user User
	db.Select("id,username").Where("username = ?", name).First(&user)
	if user.ID == uint(id) {
		return errmsg.SUCCESS
	}
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	data.Password = ScryptPW(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetUser(id int) (User, int) {
	var user User
	err := db.Model(&User{}).Select("id,username,role").Where("ID = ?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

// GetUsers 查询用户列表
func GetUsers(username string, pageSize, pageNum int) ([]User, int64) {
	//var user User
	var users []User
	var total int64
	var err error

	if username == "" {
		err = db.Model(&User{}).Select("id,username,role").Count(&total).Find(&users).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
		return users, total
	}
	err = db.Model(&User{}).Select("id,username,role").Where("username LIKE ?", username+"%").Find(&users).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error

	if err != nil {
		return nil, 0
	}
	//err = db.Model(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

// EditUser 编辑用户信息
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.UserName
	maps["role"] = data.Role
	err := db.Model(user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	err := db.Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// ScryptPW 密码加密
func ScryptPW(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{11, 22, 33, 44, 55, 66, 77, 88}
	HashPW, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPW)
	return fpw

}

// CheckLogin 登陆验证
func CheckLogin(username, password string) int {
	var user User

	db.Where("username = ?", username).First(&user)

	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}

	if ScryptPW(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}

	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RINGHT
	}
	return errmsg.SUCCESS
}
