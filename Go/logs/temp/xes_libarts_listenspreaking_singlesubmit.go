package temp

import (
	"time"
)

type XesLibartsListenspreakingSinglesubmit struct {
	Id            int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Stuid         int       `xorm:"not null default 0 comment('学生id') index INT(11)"`
	Paperid       int       `xorm:"not null default 0 comment('试卷/素材id') index INT(11)"`
	Papertype     int       `xorm:"not null default 0 comment('试卷类型（1 => 听后选择, 2 => 听后回答, 3 => 听后记录并转述, 4 => 短文朗读, 5 => 实战模拟）') index TINYINT(3)"`
	Questiontype  int       `xorm:"not null default 0 comment('题型（1 => 听后选择, 2 => 听后回答, 3 => 听后记录并转述, 4 => 短文朗读）') index TINYINT(3)"`
	Materialid    int       `xorm:"not null default 0 comment('材料id') INT(11)"`
	Answer        string    `xorm:"not null comment('学生做答数据(选择题/填空题)') TEXT"`
	Speechscore   string    `xorm:"not null comment('语音测评得分') TEXT"`
	Message       string    `xorm:"not null comment('语音测评') TEXT"`
	Url           string    `xorm:"not null comment('用户作答音频') TEXT"`
	Questionscore string    `xorm:"not null comment('用户材料得分详情') TEXT"`
	Isright       string    `xorm:"not null comment('用户作答正确情况（0 => 不关心正确情况[语音题]，1 => 正确，2 =>错误）') TEXT"`
	CreatedAt     time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt     time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('更新时间') DATETIME"`
}
