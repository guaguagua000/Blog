package model

import "database/sql"

/**
用户表结构体
*/
type User struct {
	id       int64         `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	username string        `gorm:"Type:varchar(64);NOT NULL"`
	password string        `gorm:"Type:varchar(64);NOT NULL"`
	nickname string        `gorm:"Type:varchar(64);NOT NULL"`
	avatar   string        `gorm:"Type:varchar(64);NOT NULL;DEFAULT:'public/img/default_avatar.png'"`
	sex      sql.NullInt32 `gorm:"Type:INT(1);"`
}
