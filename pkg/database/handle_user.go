package database

import (
	"errors"

	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
)

// User 用户信息
type User struct {
	UserID   string `gorm:"primary_key"`
	Name     string `gorm:"size:255;not null"`
	Tel      string `gorm:"not null;unique_index"`
	Password string `gorm:"not null"`
	Position string `gorm:"not null"`
}

// // Role 用户角色权限
// type Role struct {
// 	UserID string `gorm:"primary_key"`
// 	Admin  bool
// 	CM     bool
// 	PM     bool
// }

// // RoleType 角色权限枚举
// type RoleType int

// const (
// 	// Admin 管理员
// 	Admin RoleType = iota
// 	// CM 保留管理权限
// 	CM
// 	// PM 保留管理权限
// 	PM
// )

// VerifyUser 验证用户是否存在、密码是否正确
func VerifyUser(name string, password string) (*User, error) {
	logrus.Info("待验证用户:", name)
	logrus.Info("待验证密码:", password)

	var user User

	// 从 users 表中根据 name 字段查询数据，并将查询结果写入到 user 变量中
	if err := db.Table("users").Where("name = ?", name).First(&user).Error; err != nil {
		return nil, err
	}

	// if db.Table("users").Where("name = ?", name).First(&user).RecordNotFound() {
	// 	return nil, errors.New("账号不存在")
	// }

	// 对比数据库中的密码与登录密码是否一致
	if user.Password != password {
		return nil, errors.New("密码错误")
	}

	return &user, nil
}

// createAdminUser 创建 admin 用户
func createAdminUser() error {
	var count int64
	// 查询 users 表中是否存在 admin 用户
	if err := db.Table("users").Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		user := User{
			UserID:   xid.New().String(),
			Name:     "admin",
			Tel:      "12345678901",
			Password: "admin",
			Position: "管理员",
		}

		tx := db.Begin()

		err := tx.Table("users").Create(user).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		// err = tx.Table("roles").Create(&Role{
		// 	UserID: user.UserID,
		// }).Error
		// if err != nil {
		// 	tx.Rollback()
		// 	return err
		// }

		// err = enableUserRole(tx, user.UserID, Admin)
		// if err != nil {
		// 	tx.Rollback()
		// 	return err
		// }

		tx.Commit()
	}
	return nil
}
