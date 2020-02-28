package temp

import (
	"time"
)

type XesLibartsCloudplatformsdklog struct {
	Id         int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Subjectsdk string    `xorm:"comment('科目SDK配置(1:文科,2:理科,3:英语,4:讲座,5:体验课)') TEXT"`
	Clientsdk  string    `xorm:"comment('客户端SDK配置(1:PC端,2:iOS端,3:安卓端)') TEXT"`
	Studentsdk string    `xorm:"comment('学生SDK配置') TEXT"`
	Classsdk   string    `xorm:"comment('班级SDK配置') TEXT"`
	Plansdk    string    `xorm:"comment('场次SDK配置') TEXT"`
	Allsdk     string    `xorm:"comment('SDK全体配置') TEXT"`
	CreatedAt  time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
}
