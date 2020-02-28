package temp

import (
	"time"
)

type XesLibartsStuPicbookCount struct {
	Id             int       `xorm:"not null pk autoincr comment('自增ID') INT(11)"`
	Stuid          int       `xorm:"not null default 0 comment('学生ID') index INT(11)"`
	Readedcount    int       `xorm:"not null default 0 comment('阅读绘本的总次数') INT(11)"`
	Bookscount     int       `xorm:"not null default 0 comment('阅读的总绘本数') INT(11)"`
	Speaktimecount int       `xorm:"not null default 0 comment('阅读时间统计') INT(11)"`
	CreatedAt      time.Time `xorm:"not null default '0000-00-00 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt      time.Time `xorm:"not null default '0000-00-00 00:00:00' comment('修改时间') DATETIME"`
}
