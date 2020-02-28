package temp

import (
	"time"
)

type XesLibartsGameClothChange struct {
	Id        int       `xorm:"not null pk autoincr comment('自增id') INT(10)"`
	Stuid     int       `xorm:"not null default 0 comment('学生id') index INT(10)"`
	Clothdata string    `xorm:"not null comment('装扮数据json存储') TEXT"`
	CreatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('更新时间') DATETIME"`
}
