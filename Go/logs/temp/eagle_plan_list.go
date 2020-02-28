package temp

import (
	"time"
)

type EaglePlanList struct {
	Id           int       `xorm:"not null pk autoincr comment('主键id') INT(11)"`
	PlanId       int       `xorm:"not null default 0 comment('场次id') index INT(11)"`
	PlanName     string    `xorm:"not null default '' comment('场次名字') VARCHAR(255)"`
	TeacherId    int       `xorm:"not null default 0 comment('主讲id') index INT(11)"`
	TeacherName  string    `xorm:"not null default '' comment('主讲名字') index VARCHAR(64)"`
	CourseId     int       `xorm:"not null default 0 comment('课程id') index INT(11)"`
	CourseName   string    `xorm:"not null default '' comment('课程名字') VARCHAR(255)"`
	STime        string    `xorm:"not null default '' comment('开始时间') index VARCHAR(20)"`
	ETime        string    `xorm:"not null default '' comment('结束时间') index VARCHAR(20)"`
	LiveCodeType int       `xorm:"not null default -1 comment('直播代码类型--(2 文,0 理, 1英, 3大班, -1 未知)') index TINYINT(4)"`
	IsTest       int       `xorm:"not null default 0 comment('是否是测试场次,0 不是,1是') index TINYINT(4)"`
	Expect       int       `xorm:"not null default 0 comment('应到人数') INT(11)"`
	PlanStatus   int       `xorm:"not null default 0 comment('场次状态（0未开始，1进行中， 2已结束）') INT(11)"`
	PlanClassNum int       `xorm:"not null default 0 comment('班级数') INT(11)"`
	Grade        string    `xorm:"not null default '' comment('年级') VARCHAR(255)"`
	Gradeids     string    `xorm:"not null default '' comment('年级ids') VARCHAR(100)"`
	Subjectids   string    `xorm:"not null default '' comment('学科ids') VARCHAR(100)"`
	Subject      string    `xorm:"not null default '' comment('真实学科') VARCHAR(255)"`
	Day          string    `xorm:"not null default '' comment('场次日期, 例如:20200107') index VARCHAR(20)"`
	CreateTime   time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
}
