package temp

import (
	"time"
)

type XesLibartsRecitereadTwoLogs struct {
	Id        int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Stuid     int       `xorm:"not null default 0 comment('学生ID') index INT(11)"`
	Taskid    int       `xorm:"not null default 0 comment('任务ID') index INT(11)"`
	Score     int       `xorm:"not null default 0 comment('背诵分数') INT(11)"`
	Content   string    `xorm:"comment('结果提交') TEXT"`
	History   string    `xorm:"comment('历史提交记录') TEXT"`
	CreatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('修改时间') DATETIME"`
}
