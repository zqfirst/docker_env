package temp

import (
	"time"
)

type XesLibartsListenspreakingPaperstats struct {
	Id          int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Stuid       int       `xorm:"not null default 0 comment('学生id') index INT(11)"`
	Paperid     int       `xorm:"not null default 0 comment('试卷/素材id') index INT(11)"`
	Papertype   int       `xorm:"not null default 0 comment('试卷类型（1 => 听后选择, 2 => 听后回答, 3 => 听后记录并转述, 4 => 短文朗读, 5 => 实战模拟）') index TINYINT(3)"`
	Isusing     int       `xorm:"not null default 0 comment('试卷解锁状态(0 => 未解锁, 1 => 已解锁, 2 => 永久有效)') TINYINT(3)"`
	Expired     int       `xorm:"not null default 0 comment('试卷是否过期(0 => 未过期, 1 => 已过期)') TINYINT(3)"`
	Unlocktime  time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('解锁时间') DATETIME"`
	Overtime    time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('有效期截至时间') DATETIME"`
	Lastscore   float64   `xorm:"not null default 0.0 comment('试卷得分') DOUBLE(11,1)"`
	Totalscore  int       `xorm:"not null default 0 comment('试卷总分') INT(11)"`
	Iscompleted int       `xorm:"not null default 0 comment('试卷是否已作答（0 =>未作答，1=>已作答）') TINYINT(3)"`
	CreatedAt   time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt   time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('更新时间') DATETIME"`
}
