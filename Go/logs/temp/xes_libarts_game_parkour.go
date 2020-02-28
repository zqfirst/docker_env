package temp

import (
	"time"
)

type XesLibartsGameParkour struct {
	Id         int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Stuid      int       `xorm:"not null default 0 comment('学生id') index INT(11)"`
	Score      string    `xorm:"not null comment('本次得分') TEXT"`
	Highscore  string    `xorm:"not null comment('历史最高分') TEXT"`
	Recordtime string    `xorm:"not null comment('刷新历史高分时间') TEXT"`
	CreatedAt  time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt  time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('更新时间') DATETIME"`
}
