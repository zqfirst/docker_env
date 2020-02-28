package temp

import (
	"time"
)

type XesLibartsPicbookFiles struct {
	Id          int       `xorm:"not null pk autoincr comment('素材ID') INT(11)"`
	CatName     string    `xorm:"not null default '0' comment('类别名称') VARCHAR(20)"`
	Bid         int       `xorm:"not null default 0 comment('书籍ID') index INT(11)"`
	Cid         int       `xorm:"not null default 0 comment('类别ID') index INT(11)"`
	BookName    string    `xorm:"not null default '0' comment('书籍名称') VARCHAR(100)"`
	Mession     int       `xorm:"not null default 0 comment('mession') INT(11)"`
	WordsNum    int       `xorm:"not null default 0 comment('字数') INT(11)"`
	Level       int       `xorm:"not null default 0 comment('层次') TINYINT(4)"`
	Pic         string    `xorm:"not null default '0' comment('书籍图片路径') VARCHAR(200)"`
	Zipfile     string    `xorm:"not null default '0' comment('书籍zip文件路径') VARCHAR(200)"`
	Stage       string    `xorm:"not null default '0' comment('阶段') VARCHAR(50)"`
	Orientation int       `xorm:"not null default 0 comment('方向（横向或纵向）') TINYINT(4)"`
	Bookpages   string    `xorm:"not null comment('书籍页面内容') TEXT"`
	Quiz        string    `xorm:"not null comment('测试题目内容') TEXT"`
	CreatedAt   int       `xorm:"not null default 0 comment('创建时间') INT(11)"`
	UpdatedAt   int       `xorm:"not null default 0 comment('更新时间') INT(11)"`
	ClassType   int       `xorm:"not null default 0 comment('类别类型') TINYINT(4)"`
	IsShow      int       `xorm:"not null default 0 comment('是否展示（0：否，1：是）') TINYINT(4)"`
	PageNum     int       `xorm:"not null default 0 comment('页面数量') INT(11)"`
	QuizNum     int       `xorm:"not null default 0 comment('测试题目数量') INT(11)"`
	LeapStage   string    `xorm:"not null default '0' comment('跳过的阶段') VARCHAR(50)"`
	VisitedNums int       `xorm:"not null default 0 comment('访问数量') INT(11)"`
	DeletedAt   time.Time `xorm:"not null default '1970-01-01 00:00:00' comment('删除时间') DATETIME"`
}
