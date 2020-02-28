package temp

import (
	"time"
)

type XesLibartsVideorecordSubjects struct {
	Id         int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Sname      string    `xorm:"not null default '' comment('主题名称') VARCHAR(256)"`
	ImgUrl     string    `xorm:"not null default '' comment('主题图片链接') VARCHAR(256)"`
	ResUrl     string    `xorm:"not null default '' comment('主题资源包链接') VARCHAR(256)"`
	Status     int       `xorm:"not null default 1 comment('是否被删除：0，无效  1，正常有效') INT(8)"`
	Words      string    `xorm:"not null comment('主题导语描述') TEXT"`
	Notes      string    `xorm:"not null comment('备注') TEXT"`
	Stat       int       `xorm:"not null default 1 comment('状态') INT(4)"`
	CreateTime time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
}
