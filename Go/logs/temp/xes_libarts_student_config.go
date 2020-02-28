package temp

import (
	"time"
)

type XesLibartsStudentConfig struct {
	Id         int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Stuid      int       `xorm:"not null default 0 comment('学生id') index INT(11)"`
	Subjectid  int       `xorm:"not null default 1 comment('学科 1=语文,3=英语') index TINYINT(1)"`
	Prepconfig string    `xorm:"not null comment('用户预习配置 json 存储') TEXT"`
	CreatedAt  time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt  time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
}
