package temp

import (
	"time"
)

type XesLibartsRecordTasks struct {
	Id          int       `xorm:"not null pk autoincr INT(11)"`
	Stuid       int       `xorm:"not null default 0 comment('学生id') index index(index_stu_day) INT(11)"`
	Stucouid    int       `xorm:"not null default 0 comment('学生购课id') INT(11)"`
	Courseid    int       `xorm:"not null default 0 comment('课程id') INT(11)"`
	Total       int       `xorm:"not null default 0 comment('总讲次数') TINYINT(4)"`
	Planid      int       `xorm:"not null default 0 comment('场次id') INT(11)"`
	Planday     int       `xorm:"not null default 0 comment('场次日期') index(index_stu_day) INT(11)"`
	Planname    string    `xorm:"not null default '' comment('场次名称') VARCHAR(100)"`
	Teachername string    `xorm:"not null default '' comment('教师名称') VARCHAR(50)"`
	Teacherurl  string    `xorm:"not null default '' comment('教师头像') VARCHAR(200)"`
	Publish     int       `xorm:"not null default 0 comment('是否发布 1已发布 2未发布') TINYINT(1)"`
	Status      int       `xorm:"not null default 1 comment('1正常 2已删除') TINYINT(1)"`
	CreatedAt   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	UpdatedAt   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
}
