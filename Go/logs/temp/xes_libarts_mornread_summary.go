package temp

import (
	"time"
)

type XesLibartsMornreadSummary struct {
	Id          int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Stuid       int       `xorm:"not null default 0 comment('学生id') index INT(11)"`
	Courseid    int       `xorm:"not null default 0 comment('课程id') index INT(11)"`
	Planid      int       `xorm:"not null default 0 comment('场次id') index INT(11)"`
	Taskid      int       `xorm:"not null default 0 comment('作业(任务)id') index INT(11)"`
	Listening   int       `xorm:"not null default 0 comment('听力时长') INT(11)"`
	Repeatafter int       `xorm:"not null default 0 comment('跟读次数') INT(11)"`
	Recitation  int       `xorm:"not null default 0 comment('背诵次数') INT(11)"`
	CreatedAt   time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt   time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('更新时间') DATETIME"`
}
