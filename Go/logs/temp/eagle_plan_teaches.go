package temp

import (
	"time"
)

type EaglePlanTeaches struct {
	Id          int       `xorm:"not null pk autoincr INT(11)"`
	PlanId      int       `xorm:"not null comment('场次id') index INT(11)"`
	PlanName    string    `xorm:"not null comment('场次名字') VARCHAR(512)"`
	TeacherId   int       `xorm:"not null comment('老师id') index INT(11)"`
	TeacherName string    `xorm:"not null comment('教师名字') index VARCHAR(64)"`
	Nickname    string    `xorm:"not null comment('教师昵称') VARCHAR(64)"`
	Sex         int       `xorm:"not null comment('性别， 1：男、2：女、3：未知') TINYINT(2)"`
	Type        int       `xorm:"not null comment('教师类型 1:全职教师,2:兼职教师,3:专家,4:辅导教师') index TINYINT(4)"`
	ClassId     int       `xorm:"not null comment('班级id. 主讲班级id是0') index INT(11)"`
	CreateTime  time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
}
