package ctrl

import (
	"net/http"
	"fmt"
	"math/rand"
	"../util"
	"../service"
	"../model"
)

var userService service.UserService


//登录api
func UserLogin(writer http.ResponseWriter,
	request *http.Request) {
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")
    user,err := userService.Login(mobile,passwd)

    if err!=nil{
    	util.RespFail(writer,err.Error())
	}else{
		util.RespOk(writer,user,"")
	}

}

//注册api
func UserRegister(writer http.ResponseWriter,
	request *http.Request) {

	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	plainpwd := request.PostForm.Get("passwd")
	nickname := fmt.Sprintf("user%06d",rand.Int31())
	avatar := ""
	sex := model.SEX_UNKNOW

	user,err := userService.Register(mobile, plainpwd,nickname,avatar,sex)
	if err!=nil{
		util.RespFail(writer,err.Error())
	}else{
		util.RespOk(writer,user,"")

	}

}
