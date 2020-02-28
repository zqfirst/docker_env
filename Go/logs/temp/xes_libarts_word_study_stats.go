package temp

import (
	"time"
)

type XesLibartsWordStudyStats struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Stuid      int       `xorm:"not null default 0 comment('学生id') index(stu_course) INT(20)"`
	Courseid   int       `xorm:"not null default 0 comment('课程id') index(stu_course) INT(20)"`
	Wordcount  int       `xorm:"not null default 0 comment('用户累计学习天数') INT(10)"`
	Daycount   int       `xorm:"not null default 0 comment('用户累计学习天数') INT(10)"`
	CreatedAt  time.Time `xorm:"not null default '0000-00-00 00:00:00' comment('创建时间') DATETIME"`
	ModifyTime time.Time `xorm:"not null default '0000-00-00 00:00:00' comment('记录修改时间') DATETIME"`
	UpdatedAt  time.Time `xorm:"not null default '0000-00-00 00:00:00' comment('修改时间') DATETIME"`
}
