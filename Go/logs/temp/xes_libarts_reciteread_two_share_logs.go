package temp

import (
	"time"
)

type XesLibartsRecitereadTwoShareLogs struct {
	Id        int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Stuid     int       `xorm:"not null default 0 comment('学生ID') index INT(11)"`
	Taskid    int       `xorm:"not null default 0 comment('任务ID') INT(11)"`
	Type      int       `xorm:"not null default 0 comment('分享内容类型 结果页为1，国学页为2') INT(11)"`
	Token     string    `xorm:"not null default '0' comment('token') index VARCHAR(128)"`
	Content   string    `xorm:"comment('分享内容') TEXT"`
	CreatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('修改时间') DATETIME"`
}
