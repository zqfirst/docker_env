package xeslog

import (
	"time"
)

type Caipiao struct {
	Id       int       `xorm:"not null pk autoincr INT(10)"`
	Type     string    `xorm:"not null default '0' comment('1双色球2大乐透') VARCHAR(10)"`
	Num1     string    `xorm:"not null default '0' VARCHAR(10)"`
	Num2     string    `xorm:"not null default '0' VARCHAR(10)"`
	Num3     string    `xorm:"not null default '0' VARCHAR(10)"`
	Num4     string    `xorm:"not null default '0' VARCHAR(10)"`
	Num5     string    `xorm:"not null default '0' VARCHAR(10)"`
	Num6     string    `xorm:"not null default '0' VARCHAR(10)"`
	Num7     string    `xorm:"not null default '0' VARCHAR(10)"`
	TermDate time.Time `xorm:"DATE"`
	Term     string    `xorm:"not null default '' VARCHAR(20)"`
	CreateAt time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	UpdateAt time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
}
