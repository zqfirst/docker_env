package temp

import (
	"time"
)

type XesLibartsPreparationGame struct {
	Id        int       `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Stuid     int       `xorm:"not null default 0 comment('学生id') index INT(11)"`
	Game      int       `xorm:"not null default 0 comment('游戏 1=找单词|2=英语换装|3=跑酷|4=赛车|5=垃圾分类|6=盖楼|7=语文换装') index TINYINT(3)"`
	Score     int       `xorm:"not null default 0 comment('得分') INT(11)"`
	Highscore int       `xorm:"not null default 0 comment('历史最高分') INT(11)"`
	Extdata   string    `xorm:"not null comment('扩展数据') TEXT"`
	CreatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('创建时间') DATETIME"`
	UpdatedAt time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('更新时间') DATETIME"`
}
