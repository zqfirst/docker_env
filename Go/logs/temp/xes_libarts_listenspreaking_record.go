package temp

import (
	"time"
)

type XesLibartsListenspreakingRecord struct {
	Id          int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Stuid       int       `xorm:"not null default 0 comment('学生id') index INT(11)"`
	Paperid     int       `xorm:"not null default 0 comment('试卷/素材id') index INT(11)"`
	Papertype   int       `xorm:"not null default 0 comment('试卷类型（1 => 听后选择, 2 => 听后回答, 3 => 听后记录并转述, 4 => 短文朗读, 5 => 实战模拟）') index TINYINT(3)"`
	Singlescore string    `xorm:"not null comment('单项得分(实战模拟)') TEXT"`
	Totalscore  string    `xorm:"not null comment('总体得分') TEXT"`
	Count       int       `xorm:"not null default 0 comment('完成次数(实战模拟)') INT(11)"`
	Studytime   int       `xorm:"not null default 0 comment('学习时长(实战模拟)') INT(11)"`
	IsComplete  int       `xorm:"not null default 0 comment('是否完成当前试卷(0 => 未完成, 1 => 已完成)') TINYINT(3)"`
	CreatedAt   time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt   time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
}
