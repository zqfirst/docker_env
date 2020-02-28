package temp

import (
	"time"
)

type XesLibartsBeautyReadSignLog struct {
	Id          int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Stuid       int       `xorm:"not null default 0 comment('学生id') index INT(11)"`
	Year        int       `xorm:"not null default 0 comment('签到年份') INT(11)"`
	Month       int       `xorm:"not null default 0 comment('签到月份') INT(11)"`
	Conseqcount int       `xorm:"not null default 0 comment('至今天该课程下连续签到次数') INT(11)"`
	Detail      string    `xorm:"not null comment('签到详情') TEXT"`
	Lasttime    string    `xorm:"not null default '' comment('最近一次签到日期') VARCHAR(128)"`
	CreatedAt   time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt   time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
}
