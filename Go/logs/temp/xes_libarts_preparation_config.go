package temp

import (
	"time"
)

type XesLibartsPreparationConfig struct {
	Id        int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Stuid     int       `xorm:"not null default 0 comment('学生id') index INT(11)"`
	Subject   int       `xorm:"not null default 1 comment('学科 1=语文,3=英语') index TINYINT(1)"`
	Sound     int       `xorm:"not null default 1 comment('音效开启 1=开启,0=关闭') TINYINT(1)"`
	Music     int       `xorm:"not null default 1 comment('背景音乐开启 1=开启,0=关闭') TINYINT(1)"`
	CreatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
}
