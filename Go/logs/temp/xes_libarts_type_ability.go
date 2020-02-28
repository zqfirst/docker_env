package temp

import (
	"time"
)

type XesLibartsTypeAbility struct {
	Id           int       `xorm:"not null pk autoincr comment('自增ID') INT(11)"`
	StuId        int       `xorm:"not null default 0 comment('学生ID') index INT(11)"`
	CourseId     int       `xorm:"not null default 0 comment('课程ID') INT(11)"`
	PlanId       int       `xorm:"not null default 0 comment('场次ID') INT(11)"`
	ClassId      int       `xorm:"not null default 0 comment('班级ID') index INT(11)"`
	WorkId       int       `xorm:"not null default 0 comment('作业ID') INT(11)"`
	TypeAbility  string    `xorm:"not null comment('能力模型和题类型') TEXT"`
	LevelAbility string    `xorm:"not null comment('同一level下平均能力模型') TEXT"`
	ClassType    string    `xorm:"not null comment('班级平均题型') TEXT"`
	Level        int       `xorm:"not null default 0 comment('级别（1到14代表14个学段）') INT(11)"`
	Evaluation   string    `xorm:"not null comment('评价（好中差  3：好，2：中，1：差）') TEXT"`
	CreatedAt    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	UpdatedAt    time.Time `xorm:"not null default '0000-00-00 00:00:00' comment('更新时间') DATETIME"`
}
