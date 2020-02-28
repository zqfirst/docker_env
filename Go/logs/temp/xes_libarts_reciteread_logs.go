package temp

import (
	"time"
)

type XesLibartsRecitereadLogs struct {
	Id        int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Stuid     int       `xorm:"not null default 0 comment('学生id') index INT(11)"`
	Taskid    int       `xorm:"not null default 0 comment('视频任务id') index INT(11)"`
	Score     int       `xorm:"not null default 0 comment('用户最近一次得分') INT(11)"`
	Speaktime int       `xorm:"not null default 0 comment('累计开口时长') INT(11)"`
	Url       string    `xorm:"not null default '' comment('学生视频链接') VARCHAR(256)"`
	CreatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('更新时间') DATETIME"`
}
