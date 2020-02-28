package temp

import (
	"time"
)

type XesLibartsWordrace struct {
	Id        int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Stucouid  int       `xorm:"not null default 0 comment('学生购课id') index INT(11)"`
	Courseid  int       `xorm:"not null default 0 comment('课程id') INT(11)"`
	Planid    int       `xorm:"not null default 0 comment('场次id') index INT(11)"`
	Stuid     int       `xorm:"not null default 0 comment('学生id') index INT(11)"`
	Detail    string    `xorm:"not null comment('完成记录') TEXT"`
	CreatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('更新时间') DATETIME"`
}
