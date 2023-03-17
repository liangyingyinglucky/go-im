package model

import "time"

const (
	SEX_WOMEN="W"
	SEX_MEN="M"
	SEX_UNKNOW="U"
)

//用户登录
type User struct {
	//用户ID
	Id         int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	//手机号
	Mobile   string 		`xorm:"varchar(20)" form:"mobile" json:"mobile"`
	//密码
	Passwd       string	`xorm:"varchar(40)" form:"passwd" json:"-"`
	//头像
	Avatar	   string 		`xorm:"varchar(150)" form:"avatar" json:"avatar"`
	//性别
	Sex        string	`xorm:"varchar(2)" form:"sex" json:"sex"`
	//昵称
	Nickname    string	`xorm:"varchar(20)" form:"nickname" json:"nickname"`
	//加盐随机字符串6
	Salt       string	`xorm:"varchar(10)" form:"salt" json:"-"`
	Online     int	`xorm:"int(10)" form:"online" json:"online"`
	//前端鉴权
	Token      string	`xorm:"varchar(40)" form:"token" json:"token"`
	Memo      string	`xorm:"varchar(140)" form:"memo" json:"memo"`
	Createat   time.Time	`xorm:"datetime" form:"createat" json:"createat"`
}