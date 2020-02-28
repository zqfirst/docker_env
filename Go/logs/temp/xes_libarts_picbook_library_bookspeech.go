package temp

import (
	"time"
)

type XesLibartsPicbookLibraryBookspeech struct {
	Id        int       `xorm:"not null pk autoincr comment('自增ID') INT(11)"`
	Stuid     int       `xorm:"not null default 0 comment('学生id') index INT(11)"`
	Bookid    int       `xorm:"not null default 0 comment('绘本id') index INT(11)"`
	Pagecount int       `xorm:"not null default 0 comment('总页数') INT(11)"`
	Score     int       `xorm:"not null default 0 comment('得分数') INT(11)"`
	Speaktime int       `xorm:"not null default 0 comment('开口时长') INT(11)"`
	Star      int       `xorm:"not null default 0 comment('获得的星星数量') INT(11)"`
	CreatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('修改时间') DATETIME"`
}
