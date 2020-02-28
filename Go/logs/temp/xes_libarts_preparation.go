package temp

import (
	"time"
)

type XesLibartsPreparation struct {
	Id        int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Type      int       `xorm:"not null default 3 comment('学科 1=语文|3=英语') TINYINT(3)"`
	Stucouid  int       `xorm:"not null default 0 comment('学生购课id') index INT(11)"`
	Courseid  int       `xorm:"not null default 0 comment('课程id') INT(11)"`
	Catalogid int       `xorm:"not null default 0 comment('大纲目录id') INT(11)"`
	Planid    int       `xorm:"not null default 0 comment('场次id') index INT(11)"`
	Stuid     int       `xorm:"not null default 0 comment('学生id') index INT(11)"`
	Energy    int       `xorm:"not null default 0 comment('获得的能量条数') INT(11)"`
	Detail    string    `xorm:"not null comment('能量条获得详情') TEXT"`
	Card      int       `xorm:"not null default 0 comment('是否获得了碎片') INT(11)"`
	CreatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('更新时间') DATETIME"`
}
