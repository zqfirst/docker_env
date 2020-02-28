package temp

import (
	"time"
)

type XesLibartsStuPreparationStatus struct {
	Id             int       `xorm:"not null pk autoincr INT(11)"`
	PlanId         int       `xorm:"default 0 comment('场次id') INT(11)"`
	StuId          int       `xorm:"default 0 comment('用户id') INT(11)"`
	CompleteStatus int       `xorm:"default 0 comment('用户完成状态') TINYINT(4)"`
	CreatedAt      time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	CourseId       int       `xorm:"not null default 0 comment('课程id') INT(11)"`
}
