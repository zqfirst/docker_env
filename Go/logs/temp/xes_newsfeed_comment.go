package temp

import (
	"time"
)

type XesNewsfeedComment struct {
	Id             int64     `xorm:"pk autoincr comment('自增ID') BIGINT(20)"`
	CommentId      string    `xorm:"not null default '0' comment('评论消息id') unique VARCHAR(64)"`
	UserId         int64     `xorm:"not null default 0 comment('发送用户id') BIGINT(20)"`
	UserType       int       `xorm:"not null default 0 comment('用户类型，1官方账号 2运营管理账号 3学生账号 4主讲老师') TINYINT(3)"`
	TargetUserId   int64     `xorm:"not null default 0 comment('发送目标用户id') index BIGINT(20)"`
	TargetUserType int       `xorm:"not null default 0 comment('用户类型，1官方账号 2运营管理账号 3学生账号 4主讲老师') TINYINT(3)"`
	NewsId         string    `xorm:"not null default '' comment('评论新鲜事id') index VARCHAR(64)"`
	CommentParent  string    `xorm:"not null default '' comment('评论父级ID') index VARCHAR(64)"`
	CommentType    int       `xorm:"not null default 1 comment('评论类型，1新鲜事评论 2评论回复') TINYINT(3)"`
	CommentText    string    `xorm:"not null default '' comment('评论文字') VARCHAR(2000)"`
	Status         int       `xorm:"not null default 1 comment('状态，0：已删除   1：待审核 2:已通过 3:已拒绝') TINYINT(3)"`
	CreateTime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') index DATETIME"`
	DeleteTime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('删除时间') DATETIME"`
	DeleteBy       string    `xorm:"not null default '' comment('删除操作信息') VARCHAR(100)"`
}
