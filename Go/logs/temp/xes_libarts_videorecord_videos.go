package temp

import (
	"time"
)

type XesLibartsVideorecordVideos struct {
	Id        int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Stucouid  int       `xorm:"not null default 0 comment('学生购课id') index INT(11)"`
	Courseid  int       `xorm:"not null default 0 comment('课程id') INT(11)"`
	Planid    int       `xorm:"not null default 0 comment('场次id') index INT(11)"`
	Classid   int       `xorm:"not null default 0 comment('班级id') INT(11)"`
	Stuid     int       `xorm:"not null default 0 comment('学生id') index INT(11)"`
	Taskid    int       `xorm:"not null default 0 comment('视频任务id') index INT(11)"`
	Subjectid int       `xorm:"not null default 0 comment('视频任务id') INT(11)"`
	Thumbup   int       `xorm:"not null default 0 comment('点赞数') INT(11)"`
	Detail    string    `xorm:"not null comment('被点赞详情') TEXT"`
	ImgUrl    string    `xorm:"not null default '' comment('学生视频缩略图链接') VARCHAR(256)"`
	VideoUrl  string    `xorm:"not null default '' comment('学生视频链接') VARCHAR(256)"`
	Testi     int       `xorm:"not null default 1 comment('是否被删除：0，无效  1，正常有效') INT(8)"`
	Status    int       `xorm:"not null default 1 comment('是否被删除：0，无效  1，正常有效') INT(8)"`
	CreatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
}
