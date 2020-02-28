package temp

import (
	"time"
)

type XesLibartsPicbookSinglespeech struct {
	Id        int       `xorm:"not null pk autoincr comment('自增ID') INT(11)"`
	Courseid  int       `xorm:"not null default 0 comment('课程ID') index INT(11)"`
	Planid    int       `xorm:"not null default 0 comment('场次ID') index INT(11)"`
	Stuid     int       `xorm:"not null default 0 comment('学生ID') index INT(11)"`
	Bookid    int       `xorm:"not null default 0 comment('绘本ID') index INT(11)"`
	Page      int       `xorm:"not null default 0 comment('书籍页') INT(11)"`
	Pagecount int       `xorm:"not null default 0 comment('书籍总页数') INT(11)"`
	Url       string    `xorm:"not null default '' comment('单题音频（最近一次）') VARCHAR(255)"`
	Score     int       `xorm:"not null default 0 comment('最近一次分数') INT(11)"`
	Speaktime int       `xorm:"not null default 0 comment('开口时长') INT(11)"`
	CreatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('修改时间') DATETIME"`
}
