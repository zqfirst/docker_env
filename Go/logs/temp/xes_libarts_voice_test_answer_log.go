package temp

import (
	"time"
)

type XesLibartsVoiceTestAnswerLog struct {
	Id         int       `xorm:"not null pk autoincr INT(10)"`
	StuId      int       `xorm:"not null default 0 comment('用户ID') INT(11)"`
	LiveId     int       `xorm:"INT(11)"`
	TestId     int       `xorm:"not null default 0 comment('试题ID') INT(11)"`
	Gold       int       `xorm:"not null default 0 comment('金币') INT(11)"`
	SrcType    int       `xorm:"not null default 0 comment('试题来源，1：普互，2：课件互动') INT(11)"`
	IsRight    int       `xorm:"not null default 0 comment('是否正确') TINYINT(4)"`
	VoiceTime  string    `xorm:"not null default '0' comment('开口时长') VARCHAR(11)"`
	VoiceUrl   string    `xorm:"not null default '' comment('语音链接') VARCHAR(128)"`
	CreateTime time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	ClassId    int       `xorm:"default 0 comment('用户班级id') INT(11)"`
}
