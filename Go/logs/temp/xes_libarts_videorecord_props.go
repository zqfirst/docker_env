package temp

import (
	"time"
)

type XesLibartsVideorecordProps struct {
	Id         int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Pname      string    `xorm:"not null default '' comment('主题名称') VARCHAR(256)"`
	PropUrl    string    `xorm:"not null default '' comment('背景音乐音频') VARCHAR(256)"`
	Notes      string    `xorm:"not null comment('备注') TEXT"`
	Stat       int       `xorm:"not null default 1 comment('课程id') INT(4)"`
	CreateTime time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
}
