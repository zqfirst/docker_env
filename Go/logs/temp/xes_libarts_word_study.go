package temp

import (
	"time"
)

type XesLibartsWordStudy struct {
	Id            int       `xorm:"not null pk autoincr comment('自增id') INT(20)"`
	Stuid         int       `xorm:"not null default 0 comment('学生id') index INT(20)"`
	Courseid      int       `xorm:"not null default 0 comment('课程id') index INT(11)"`
	Chapterid     int       `xorm:"not null default 0 comment('章节id') INT(10)"`
	Chaptername   string    `xorm:"not null comment('章节名称') VARCHAR(255)"`
	Paperid       int       `xorm:"not null default 0 comment('试卷id') index INT(10)"`
	Errorcount    int       `xorm:"not null default 0 comment('错题个数') INT(4)"`
	Questioncount int       `xorm:"not null default 0 comment('试题总数') INT(4)"`
	Stuanswer     string    `xorm:"not null comment('学生做答数据') TEXT"`
	Paperscore    int       `xorm:"not null default 0 comment('试卷得分') INT(4)"`
	Type          int       `xorm:"not null default 0 comment('类型（1:评测2:闯关）') INT(4)"`
	CreatedAt     time.Time `xorm:"not null default '0000-00-00 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt     time.Time `xorm:"not null default '0000-00-00 00:00:00' comment('修改时间') DATETIME"`
}
