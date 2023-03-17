package ctrl


import (
	"net/http"
	"../util"
	"../service"
	"../args"
	"../model"
	)

var contactService service.ContactService
//好友
func LoadFriend(w http.ResponseWriter, req *http.Request){
	var arg args.ContactArg
	//如果这个用的上,那么可以直接
	util.Bind(req,&arg)

	users := contactService.SearchFriend(arg.Userid)
	util.RespOkList(w,users,len(users))
}

//添加好友
func Addfriend(w http.ResponseWriter, req *http.Request) {
	//定义一个参数结构体
	var arg args.ContactArg
	util.Bind(req,&arg)
	//调用service
	err := contactService.AddFriend(arg.Userid,arg.Dstid)
	if err!=nil{
		util.RespFail(w,err.Error())
	}else{
		util.RespOk(w,nil,"好友添加成功")
	}
}

//群
func LoadCommunity(w http.ResponseWriter, req *http.Request){
	var arg args.ContactArg
	util.Bind(req,&arg)
	comunitys := contactService.SearchComunity(arg.Userid)
	util.RespOkList(w,comunitys,len(comunitys))
}

//加入群
func JoinCommunity(w http.ResponseWriter, req *http.Request){
	var arg args.ContactArg
	util.Bind(req,&arg)
	err := contactService.JoinCommunity(arg.Userid,arg.Dstid);
	//刷新用户的群组信息
	AddGroupId(arg.Userid,arg.Dstid)
	if err!=nil{
		util.RespFail(w,err.Error())
	}else {
		util.RespOk(w,nil,"")
	}
}

//创建群
func CreateCommunity(w http.ResponseWriter, req *http.Request){
	var arg model.Community
	util.Bind(req,&arg)
	com,err := contactService.CreateCommunity(arg);
	if err!=nil{
		util.RespFail(w,err.Error())
	}else {
		util.RespOk(w,com,"")
	}
}
