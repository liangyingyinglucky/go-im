package model

import "time"

//群信息表
type Community struct {
	Id         int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	//名称
	Name   string 		`xorm:"varchar(30)" form:"name" json:"name"`
	//群主ID
	Ownerid       int64	`xorm:"bigint(20)" form:"ownerid" json:"ownerid"`
	//群logo
	Icon	   string 		`xorm:"varchar(250)" form:"icon" json:"icon"`
	//类型
	Categroy      int	`xorm:"int(11)" form:"categroy" json:"categroy"`
	//描述
	Memo    string	`xorm:"varchar(120)" form:"memo" json:"memo"`
	//创建时间
	Createat   time.Time	`xorm:"datetime" form:"createat" json:"createat"`
}