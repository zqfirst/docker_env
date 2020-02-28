package temp

import (
	"time"
)

type XesLibartsGuideLog struct {
	Id          int       `xorm:"not null pk autoincr comment('自增ID') INT(11)"`
	Courseid    int       `xorm:"not null default 0 comment('课程ID') index INT(11)"`
	Planid      int       `xorm:"not null default 0 comment('场次ID') index INT(11)"`
	Stuid       int       `xorm:"not null default 0 comment('学生ID') index INT(11)"`
	Stucourseid int       `xorm:"not null default 0 comment('购课ID') index INT(11)"`
	Taskid      int       `xorm:"not null default 0 comment('地图ID') INT(11)"`
	Result      string    `xorm:"not null comment('结果提交') TEXT"`
	History     string    `xorm:"not null comment('历史提交记录') TEXT"`
	IsDelete    int       `xorm:"not null default 0 comment('删除标记') INT(11)"`
	CreatedAt   time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt   time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('修改时间') DATETIME"`
}
