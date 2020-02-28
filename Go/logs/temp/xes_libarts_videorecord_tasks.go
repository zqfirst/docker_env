package temp

import (
	"time"
)

type XesLibartsVideorecordTasks struct {
	Id          int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Tname       string    `xorm:"not null default '' comment('任务名称') VARCHAR(256)"`
	TaskUrl     string    `xorm:"not null default '' comment('任务图片') VARCHAR(256)"`
	Description string    `xorm:"not null comment('任务描述信息') TEXT"`
	Stat        int       `xorm:"not null default 1 comment('状态') INT(4)"`
	CreateTime  time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdateTime  time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('更新时间') DATETIME"`
}
