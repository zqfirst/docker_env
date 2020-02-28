package temp

import (
	"time"
)

type XesLibartsCompositionCommits struct {
	Id                 int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	CourseId           int       `xorm:"not null default 0 comment('课程id') INT(11)"`
	ClassId            int       `xorm:"not null default 0 comment('班id') index(class_plan_work_id) INT(11)"`
	PlanId             int       `xorm:"not null default 0 comment('场次明细id') index(class_plan_work_id) INT(11)"`
	OutlineId          int       `xorm:"not null default 0 comment('大纲明细id') INT(11)"`
	StuCouId           int       `xorm:"not null default 0 comment('用户购课自增id') INT(11)"`
	StuId              int       `xorm:"not null default 0 comment('用户id') INT(11)"`
	WorkId             int       `xorm:"not null default 0 comment('作业id') index(class_plan_work_id) INT(11)"`
	Title              string    `xorm:"not null default '' comment('作文标题') VARCHAR(100)"`
	Content            string    `xorm:"not null default '' comment('作文内容') VARCHAR(5000)"`
	CommitTime         time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('学生提交作业时间') index(class_plan_work_id) DATETIME"`
	CorrectTime        time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('老师批改作业时间') DATETIME"`
	RejectReason       string    `xorm:"not null default '' comment('作业被拒理由') VARCHAR(500)"`
	Status             int       `xorm:"not null default 1 comment('作业状态,1:已提交,2:驳回,3:已批改') TINYINT(1)"`
	AnswerLength       string    `xorm:"not null default '' comment('作答时长') CHAR(7)"`
	WordCount          string    `xorm:"not null default '' comment('作文字数') CHAR(3)"`
	PigaiKey           string    `xorm:"not null default '' comment('批改网每次请求的key') CHAR(50)"`
	IsMachineCorrected int       `xorm:"not null default 0 comment('机器是否批改完成,0:未批改完成,1:批改完成') TINYINT(1)"`
}
