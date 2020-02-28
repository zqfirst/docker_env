package temp

import (
	"time"
)

type XesLibartsMornreadPublishes struct {
	Id          int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Planid      int       `xorm:"not null default 0 comment('场次id') index INT(11)"`
	Taskid      int       `xorm:"not null default 0 comment('作业id') index INT(11)"`
	Taskname    string    `xorm:"not null default '' comment('作业名称') VARCHAR(128)"`
	Status      int       `xorm:"not null default 1 comment('是否已经发布 1:已经发布,2:取消发布,默认已经发布') TINYINT(1)"`
	Publishtime time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('作业发布时间') DATETIME"`
	CreatedAt   time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt   time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
}
