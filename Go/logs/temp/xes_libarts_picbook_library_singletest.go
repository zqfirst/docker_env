package temp

import (
	"time"
)

type XesLibartsPicbookLibrarySingletest struct {
	Id        int       `xorm:"not null pk autoincr comment('自增ID') INT(11)"`
	Stuid     int       `xorm:"not null default 0 comment('学生ID') index INT(11)"`
	Bookid    int       `xorm:"not null default 0 comment('绘本ID') index INT(11)"`
	Quizid    int       `xorm:"not null default 0 comment('题目ID') INT(11)"`
	Quizcount int       `xorm:"not null default 0 comment('总题数') INT(11)"`
	Answer    int       `xorm:"not null default 0 comment('单题作答答案（最近一次）') INT(11)"`
	CreatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('修改时间') DATETIME"`
}
