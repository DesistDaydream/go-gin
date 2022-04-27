package database

import "gorm.io/gorm"

var (
	err error
	db  *gorm.DB
)

type DBHandler interface {
	ConnDB()
}
