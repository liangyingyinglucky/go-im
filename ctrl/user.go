package ctrl

import (
	"fmt"
	"math/rand"
	"../util"
	"../service"
	"../model"
	"../web"
)

var userService service.UserService


//登录api
func UserLogin(c *web.Context) {
	c.R.ParseForm()

	mobile := c.R.PostForm.Get("mobile")
	passwd := c.R.PostForm.Get("passwd")
    user,err := userService.Login(mobile,passwd)

    if err!=nil{
    	util.RespFail(c.W,err.Error())
	}else{
		util.RespOk(c.W,user,"")
	}

}

//注册api
func UserRegister(c *web.Context) {

	c.R.ParseForm()
	mobile := c.R.PostForm.Get("mobile")
	plainpwd := c.R.PostForm.Get("passwd")
	nickname := fmt.Sprintf("user%06d",rand.Int31())
	avatar := ""
	sex := model.SEX_UNKNOW

	user,err := userService.Register(mobile, plainpwd,nickname,avatar,sex)
	if err!=nil{
		util.RespFail(c.W,err.Error())
	}else{
		util.RespOk(c.W,user,"")

	}

}
