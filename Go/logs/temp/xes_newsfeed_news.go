package temp

import (
	"time"
)

type XesNewsfeedNews struct {
	Id         int64     `xorm:"pk autoincr comment('自增ID') BIGINT(20)"`
	UserId     int64     `xorm:"not null default 0 comment('用户id') index BIGINT(20)"`
	UserType   int       `xorm:"not null default 0 comment('用户类型，1官方账号 2运营管理账号 3学生账号 4主讲老师 ') TINYINT(3)"`
	NewsId     int64     `xorm:"not null default 0 comment('新鲜事id') unique BIGINT(20)"`
	NewsText   string    `xorm:"not null default '' comment('新鲜事文字') VARCHAR(2000)"`
	NewsPic    string    `xorm:"not null default '' comment('新鲜事图片信息') VARCHAR(5000)"`
	NewsVideo  string    `xorm:"not null default '' comment('新鲜事视频信息') VARCHAR(1000)"`
	StarInfo   string    `xorm:"not null default '' comment('点赞用户信息') VARCHAR(5000)"`
	Comments   int       `xorm:"not null default 0 comment('评论新鲜事及回复评论的已通过的总数') INT(20)"`
	Status     int       `xorm:"not null default 0 comment('状态，0正常 1标记删除') TINYINT(3)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') index DATETIME"`
	DeleteTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('删除时间') DATETIME"`
	DeleteBy   string    `xorm:"not null default '' comment('删除操作信息') VARCHAR(100)"`
}
