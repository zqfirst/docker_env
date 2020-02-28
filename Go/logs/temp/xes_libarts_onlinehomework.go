package temp

import (
	"time"
)

type XesLibartsOnlinehomework struct {
	Id                          int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Stuid                       int       `xorm:"not null default 0 comment('学生id') index INT(11)"`
	Courseid                    int       `xorm:"not null default 0 comment('课程id') index INT(11)"`
	Planid                      int       `xorm:"not null default 0 comment('场次id') index INT(11)"`
	Catalogid                   int       `xorm:"not null default 0 comment('大纲目录id') INT(11)"`
	Paperid                     int       `xorm:"not null default 0 comment('试卷id') index INT(11)"`
	Papername                   string    `xorm:"not null default '' comment('试卷名称') VARCHAR(256)"`
	Score                       int       `xorm:"not null default 0 comment('得分') INT(11)"`
	Answerdetail                string    `xorm:"not null comment('用户作答详情') TEXT"`
	Righttestnum                int       `xorm:"not null default 0 comment('用户答对个数') INT(11)"`
	Righttestids                string    `xorm:"not null comment('用户答对试题集合') TEXT"`
	Wrongtestnum                int       `xorm:"not null default 0 comment('用户答错个数') INT(11)"`
	Wrongtestids                string    `xorm:"not null comment('用户答错试题集合') TEXT"`
	Rightwrongtestdetailed      string    `xorm:"not null comment('用户答题对错明细') TEXT"`
	Rightknowledgenum           int       `xorm:"not null default 0 comment('答对知识点数') INT(11)"`
	Rightknowledgeids           string    `xorm:"not null comment('答对知识id集合') TEXT"`
	Wrongknowledgenum           int       `xorm:"not null default 0 comment('答错知识点数') INT(11)"`
	Wrongknowledgeids           string    `xorm:"not null comment('答错知识id集合') TEXT"`
	Rightwrongknowledgedetailed string    `xorm:"not null comment('作答知识点明细') TEXT"`
	Starttime                   int       `xorm:"not null default 0 comment('作答开始时间') INT(11)"`
	Endtime                     int       `xorm:"not null default 0 comment('作答结束时间') INT(11)"`
	CreatedAt                   time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt                   time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
}
