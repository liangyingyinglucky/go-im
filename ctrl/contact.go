package ctrl


import (
	"../util"
	"../service"
	"../args"
	"../model"
	"../web"
	)

var contactService service.ContactService
//好友
func LoadFriend(c *web.Context){
	var arg args.ContactArg
	//如果这个用的上,那么可以直接
	util.Bind(c.R,&arg)

	users := contactService.SearchFriend(arg.Userid)
	util.RespOkList(c.W,users,len(users))
}

//添加好友
func Addfriend(c *web.Context) {
	//定义一个参数结构体
	var arg args.ContactArg
	util.Bind(c.R,&arg)
	//调用service
	err := contactService.AddFriend(arg.Userid,arg.Dstid)
	if err!=nil{
		util.RespFail(c.W,err.Error())
	}else{
		util.RespOk(c.W,nil,"好友添加成功")
	}
}

//群
func LoadCommunity(c *web.Context){
	var arg args.ContactArg
	util.Bind(c.R,&arg)
	comunitys := contactService.SearchComunity(arg.Userid)
	util.RespOkList(c.W,comunitys,len(comunitys))
}

//加入群
func JoinCommunity(c *web.Context){
	var arg args.ContactArg
	util.Bind(c.R,&arg)
	err := contactService.JoinCommunity(arg.Userid,arg.Dstid);
	//刷新用户的群组信息
	AddGroupId(arg.Userid,arg.Dstid)
	if err!=nil{
		util.RespFail(c.W,err.Error())
	}else {
		util.RespOk(c.W,nil,"")
	}
}

//创建群
func CreateCommunity(c *web.Context){
	var arg model.Community
	util.Bind(c.R,&arg)
	com,err := contactService.CreateCommunity(arg);
	if err!=nil{
		util.RespFail(c.W,err.Error())
	}else {
		util.RespOk(c.W,com,"")
	}
}
