package temp

import (
	"time"
)

type XesLibartsMornreadSpeechevalLog struct {
	Id          int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Courseid    int       `xorm:"not null default 0 comment('课程id') index INT(11)"`
	Planid      int       `xorm:"not null default 0 comment('场次id') index INT(11)"`
	Stuid       int       `xorm:"not null default 0 comment('学生id') index INT(11)"`
	Taskid      int       `xorm:"not null default 0 comment('作业（任务）id') index INT(11)"`
	Materialid  int       `xorm:"not null default 0 comment('试题材料id') INT(11)"`
	Testid      int       `xorm:"not null default 0 comment('试题id') INT(11)"`
	Sequence    int       `xorm:"not null default 0 comment('晨读指在材料的顺序') INT(11)"`
	Score       int       `xorm:"not null default 0 comment('用户最高得分') INT(11)"`
	Speaktime   int       `xorm:"not null default 0 comment('单题累计开口时长') INT(11)"`
	Answercount int       `xorm:"not null default 0 comment('单题作答次数') INT(11)"`
	Detail      string    `xorm:"not null comment('试题详细得分') TEXT"`
	Stem        string    `xorm:"not null default '' comment('题目题干') VARCHAR(1024)"`
	Analytic    string    `xorm:"not null default '' comment('题目解析（译文等）') VARCHAR(1024)"`
	Url         string    `xorm:"not null default '' comment('用户作答音频') VARCHAR(128)"`
	CreatedAt   time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt   time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
}
