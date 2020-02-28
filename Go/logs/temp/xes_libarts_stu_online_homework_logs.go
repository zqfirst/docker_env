package temp

import (
	"time"
)

type XesLibartsStuOnlineHomeworkLogs struct {
	Id                          int       `xorm:"not null pk autoincr INT(11)"`
	StuId                       int       `xorm:"not null default 0 comment('用户id') index INT(11)"`
	StuName                     string    `xorm:"not null default '' comment('用户名') CHAR(32)"`
	CourseId                    int       `xorm:"not null default 0 comment('用户id') index INT(11)"`
	CourseName                  string    `xorm:"not null default '' comment('课程名称') VARCHAR(100)"`
	OutlineCatalogId            int       `xorm:"not null default 0 comment('课程大纲明细id') INT(11)"`
	OutlineCatalogName          string    `xorm:"not null default '' comment('课程大纲明细名称') VARCHAR(100)"`
	PaperId                     int       `xorm:"not null default 0 comment('试卷id') index INT(11)"`
	PaperName                   string    `xorm:"not null default '' comment('试卷名称') VARCHAR(100)"`
	StuScore                    int       `xorm:"not null default 0 comment('用户得分') INT(11)"`
	RightTestNum                int       `xorm:"not null default 0 comment('作答正确试题数') INT(11)"`
	RightTestIds                string    `xorm:"not null default '' comment('作答正确试题id集合(id间用,隔开)') VARCHAR(1000)"`
	WrongTestNum                int       `xorm:"not null default 0 comment('作答错误试题数') INT(11)"`
	WrongTestIds                string    `xorm:"not null default '' comment('作答错误试题id集合(id间用.隔开)') VARCHAR(1000)"`
	RightWrongTestDetailed      string    `xorm:"not null comment('用户答案对错明细') TEXT"`
	RightKnowledgeNum           int       `xorm:"not null default 0 comment('作答正确知识点个数') INT(11)"`
	RightKnowledgeIds           string    `xorm:"not null default '' comment('作答正确知识点id集合(id间用,隔开)') VARCHAR(255)"`
	WrongKnowledgeNum           int       `xorm:"not null default 0 comment('作答错误知识点数') INT(11)"`
	WrongKnowledgeIds           string    `xorm:"not null default '' comment('用户作答错误知识点id集合(id间用,隔开)') VARCHAR(255)"`
	RightWrongKnowledgeDetailed string    `xorm:"not null comment('试题知识点详细数据') TEXT"`
	StuAnswer                   string    `xorm:"not null comment('用户提交答案信息(序列化后信息)') TEXT"`
	AnswerUseTime               int       `xorm:"not null default 0 comment('作答总用时(单位：秒)') INT(11)"`
	AnswerStartTime             time.Time `xorm:"not null default '0000-00-00 00:00:00' comment('用户作答开始时间') DATETIME"`
	AnswerEndTime               time.Time `xorm:"not null default '0000-00-00 00:00:00' comment('用户作答结束时间') DATETIME"`
	CreatedAt                   time.Time `xorm:"not null default '0000-00-00 00:00:00' comment('记录创建时间') DATETIME"`
	PlanId                      int       `xorm:"not null default 0 INT(11)"`
}
