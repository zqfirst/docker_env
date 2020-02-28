package temp

import (
	"time"
)

type XesLibartsCoursedetail struct {
	Id           int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Courseid     int       `xorm:"not null default 0 comment('课程id') index INT(11)"`
	Planid       int       `xorm:"not null default 0 comment('场次id') index INT(11)"`
	Stuid        int       `xorm:"not null default 0 comment('学生id') index INT(11)"`
	Type         int       `xorm:"not null default 0 comment('任务类型id 1预习，2直播，3随堂测，4复习') INT(11)"`
	Answerstatus string    `xorm:"not null default '' comment('用户作答音频') VARCHAR(256)"`
	CreatedAt    time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt    time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
}
