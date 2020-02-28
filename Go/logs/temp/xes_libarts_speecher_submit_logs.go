package temp

import (
	"time"
)

type XesLibartsSpeecherSubmitLogs struct {
	Id            int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	CourseId      int       `xorm:"not null default 0 comment('课程id') index INT(11)"`
	PlanId        int       `xorm:"not null default 0 comment('场次id') index INT(11)"`
	PlanName      string    `xorm:"not null default '' comment('场次名称') VARCHAR(255)"`
	PlanSort      int       `xorm:"not null default 0 comment('场次排序') INT(11)"`
	ClassId       int       `xorm:"not null default 0 comment('班级id') INT(11)"`
	TaskId        int       `xorm:"not null default 0 comment('主题id') INT(11)"`
	MaterialId    int       `xorm:"not null default 0 comment('素材id') INT(11)"`
	ActivityId    int       `xorm:"not null default 0 comment('活动id') INT(11)"`
	ActivityName  string    `xorm:"not null default '' comment('活动名称') VARCHAR(255)"`
	StuId         int       `xorm:"not null default 0 comment('学生id') index INT(11)"`
	StuCourseId   int       `xorm:"not null default 0 comment('学生课程id') INT(11)"`
	AiTaskId      int       `xorm:"not null default 0 comment('AI的任务id') INT(11)"`
	AiResult      int       `xorm:"not null default 0 comment('ai返回状态0:尚未返回 1:成功 2:失败') INT(11)"`
	AiSubmitTime  int       `xorm:"not null default 0 comment('提交AI评分的时间') INT(11)"`
	ImgUrl        string    `xorm:"not null default '' comment('学生视频缩略图链接') VARCHAR(255)"`
	VideoUrl      string    `xorm:"not null default '' comment('学生视频url') VARCHAR(255)"`
	VoiceUrl      string    `xorm:"not null default '' comment('学生音频url') VARCHAR(255)"`
	SubmitTime    int       `xorm:"not null default 0 comment('提交时间') INT(11)"`
	Thumbup       int       `xorm:"not null default 0 comment('点赞数') INT(11)"`
	ScoreA        int       `xorm:"not null default 0 comment('维度a分数(语言流畅)') INT(11)"`
	ScoreB        int       `xorm:"not null default 0 comment('维度b分数(开口音量)') INT(11)"`
	ScoreC        int       `xorm:"not null default 0 comment('维度c分数(中心明确)') INT(11)"`
	ScoreD        int       `xorm:"not null default 0 comment('维度d分数(专注程度)') INT(11)"`
	TipOffNum     int       `xorm:"not null default 0 comment('举报次数') INT(11)"`
	TipOffTypeNum string    `xorm:"not null default '0,0,0' comment('每个维度举报个数') VARCHAR(255)"`
	Status        int       `xorm:"not null default 1 comment('是否被删除：0，无效 1，正常有效') INT(8)"`
	CreateTime    time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
}
