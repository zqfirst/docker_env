package temp

type XesLibartsVoiceBobVersion struct {
	Id          int    `xorm:"not null pk autoincr INT(11)"`
	Version     string `xorm:"comment('版本号') VARCHAR(20)"`
	Description string `xorm:"comment('描述') VARCHAR(128)"`
	Url         string `xorm:"comment('下载地址') VARCHAR(255)"`
	CreateTime  string `xorm:"comment('创建时间') VARCHAR(32)"`
}
