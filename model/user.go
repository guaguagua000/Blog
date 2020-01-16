package model

import "database/sql"

/**
用户表结构体
*/
type User struct {
	Id       int64         `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Username string        `gorm:"Type:varchar(64);NOT NULL"`
	Password string        `gorm:"Type:varchar(64);NOT NULL"`
	Nickname string        `gorm:"Type:varchar(64);NOT NULL"`
	Avatar   string        `gorm:"Type:varchar(64);NOT NULL;DEFAULT:'public/img/default_avatar.png'"`
	Sex      sql.NullInt32 `gorm:"Type:INT(1);"`
}
