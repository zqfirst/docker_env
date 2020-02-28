package temp

import (
	"time"
)

type XesLibartsStudentStudyTasks struct {
	Id           int       `xorm:"not null pk autoincr comment('自增ID') INT(10)"`
	Stuid        int       `xorm:"not null default 0 comment('学生id') index index(stucatalog) index(stucouplan) index(stuplantypestatus) index(stustatusetime) INT(11)"`
	Stucouid     int       `xorm:"not null default 0 comment('学生课程id') index(stucatalog) index(stucoucatalog) index(stucouplan) INT(11)"`
	Catalogid    int       `xorm:"not null default 0 comment('大纲目录id') index(stucatalog) index(stucoucatalog) INT(11)"`
	Courseid     int       `xorm:"not null default 0 comment('课程id') index(stuplantypestatus) INT(11)"`
	Classid      int       `xorm:"not null default 0 comment('直播场次id') INT(11)"`
	Planid       int       `xorm:"not null default 0 comment('直播场次id') index index(stucouplan) index(stuplantypestatus) INT(11)"`
	Originplanid int       `xorm:"not null default 0 comment('原始场次ID') INT(11)"`
	Etime        int       `xorm:"not null default 0 comment('场次结束时间戳') index index(stustatusetime) INT(10)"`
	Tasktype     int       `xorm:"not null default 0 comment('类型，1预习 2直播 3随堂测 4复习') index(stuplantypestatus) TINYINT(3)"`
	Isstd        int       `xorm:"not null default 0 comment('是否是标准任务') TINYINT(2)"`
	Isgroup      int       `xorm:"not null default 0 comment('是否是组任务') TINYINT(2)"`
	Status       int       `xorm:"not null default 1 comment('状态（默认1），1未完成 2已完成 3关闭') index(stustatusetime) TINYINT(2)"`
	Statusdetail string    `xorm:"not null default '' comment('状态详情') VARCHAR(256)"`
	CreatedAt    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	UpdatedAt    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
}
