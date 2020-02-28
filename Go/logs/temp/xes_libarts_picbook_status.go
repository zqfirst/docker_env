package temp

import (
	"time"
)

type XesLibartsPicbookStatus struct {
	Id        int       `xorm:"not null pk autoincr comment('自增ID') INT(11)"`
	Courseid  int       `xorm:"not null default 0 comment('课程ID') index INT(11)"`
	Planid    int       `xorm:"not null default 0 comment('场次ID') index INT(11)"`
	Stuid     int       `xorm:"not null default 0 comment('学生ID') index INT(11)"`
	Bookid    int       `xorm:"not null default 0 comment('绘本ID') index INT(11)"`
	Listening int       `xorm:"not null default 0 comment('听状态') INT(11)"`
	Reading   int       `xorm:"not null default 0 comment('跟读状态') INT(11)"`
	Test      int       `xorm:"not null default 0 comment('测试题状态') INT(11)"`
	CreatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('修改时间') DATETIME"`
}
