package model

import "time"

//好友和群关系
type Contact struct {
	Id         int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	//属于谁的好友
	Ownerid       int64	`xorm:"bigint(20)" form:"ownerid" json:"ownerid"`
	//好友是谁
	Dstobj       int64	`xorm:"bigint(20)" form:"dstobj" json:"dstobj"`
	//类型
	Categroy      int	`xorm:"int(11)" form:"categroy" json:"categroy"`
	//备注
	Memo    string	`xorm:"varchar(120)" form:"memo" json:"memo"`
	//创建时间
	Createat   time.Time	`xorm:"datetime" form:"createat" json:"createat"`
}

const (
		CONCAT_CATE_USER = 0x01 //好友
	    CONCAT_CATE_COMUNITY = 0x02 //群
	)