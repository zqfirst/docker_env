package temp

import (
	"time"
)

type XesLibartsGameFindSomething struct {
	Id        int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Stuid     int       `xorm:"not null default 0 comment('学生id') index INT(11)"`
	Task      string    `xorm:"not null comment('目标词组json存储') TEXT"`
	Props     string    `xorm:"not null comment('道具json存储') TEXT"`
	Words     string    `xorm:"not null comment('完成的词组json存储') TEXT"`
	Wordsnum  int       `xorm:"not null default 0 comment('完成词组的数量') TINYINT(5)"`
	Timeleft  int       `xorm:"not null default 0 comment('剩余时间,单位为秒') TINYINT(5)"`
	Isover    int       `xorm:"not null default 0 comment('游戏结束标识') index TINYINT(1)"`
	Extradata string    `xorm:"not null comment('运行数据json存储') TEXT"`
	CreatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('更新时间') DATETIME"`
}
