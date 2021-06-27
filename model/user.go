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
	UserName string `gorm:"column:username;type: varchar(20);not null" json:"username"`
	Password string `gorm:"column:password;type: varchar(20);not null" json:"password"`
	Role     int    `gorm:"column:role;type: int" json:"role"`
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

// CreateUser 新增用户
func CreateUser(data *User) int {
	data.Password = ScryptPW(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetUsers 查询用户列表
func GetUsers(pageSize, pageNum int) []User {
	var users []User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
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
