package temp

import (
	"time"
)

type XesNewsfeedUserinfo struct {
	Id         int64     `xorm:"pk autoincr comment('自增ID') BIGINT(20)"`
	UserId     int64     `xorm:"not null default 0 comment('用户id') index(uidtype) BIGINT(20)"`
	UserType   int       `xorm:"not null default 0 comment('用户类型，1官方账号,2运营管理,3学生&家长,4主讲老师') index(uidtype) TINYINT(3)"`
	UserGrade  int       `xorm:"not null default 0 comment('用户当前年级') INT(10)"`
	UserName   string    `xorm:"not null default ' ' comment('用户真实姓名') VARCHAR(20)"`
	NickName   string    `xorm:"not null default ' ' comment('用户昵称') VARCHAR(20)"`
	EnName     string    `xorm:"not null default ' ' comment('用户英文名称') VARCHAR(20)"`
	Avatar     string    `xorm:"not null default ' ' comment('用户头像链接') VARCHAR(256)"`
	Subjects   string    `xorm:"not null default ' ' comment('用户学科集合') VARCHAR(256)"`
	Grades     string    `xorm:"not null default ' ' comment('用户年级集合') VARCHAR(256)"`
	Attentions string    `xorm:"comment('用户关注id集合') TEXT"`
	CreatedAt  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	UpdatedAt  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
}
