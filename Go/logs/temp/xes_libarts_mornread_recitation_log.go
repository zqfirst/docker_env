package temp

import (
	"time"
)

type XesLibartsMornreadRecitationLog struct {
	Id          int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Stuid       int       `xorm:"not null default 0 comment('学生id') index INT(11)"`
	Courseid    int       `xorm:"not null default 0 comment('课程id') index INT(11)"`
	Planid      int       `xorm:"not null default 0 comment('场次id') index INT(11)"`
	Taskid      int       `xorm:"not null default 0 comment('作业(任务)id') index INT(11)"`
	Textid      int       `xorm:"not null default 0 comment('课文id') index INT(11)"`
	Score       int       `xorm:"not null default 0 comment('背诵得分') INT(11)"`
	Answercount int       `xorm:"not null default 0 comment('背诵次数') INT(11)"`
	Speaktime   int       `xorm:"not null default 0 comment('累计开口时长') INT(11)"`
	Hints       int       `xorm:"not null default 0 comment('用户背诵提示次数') SMALLINT(6)"`
	Audio       string    `xorm:"not null default '' comment('用户背诵音频') VARCHAR(128)"`
	Detail      string    `xorm:"not null comment('用户背诵音频') TEXT"`
	Title       string    `xorm:"not null default '' comment('背诵标题') VARCHAR(256)"`
	Body        string    `xorm:"not null comment('背诵内容') TEXT"`
	Sequence    int       `xorm:"not null default 1 comment('背诵在材料中的顺序') TINYINT(4)"`
	CreatedAt   time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt   time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('更新时间') DATETIME"`
}
