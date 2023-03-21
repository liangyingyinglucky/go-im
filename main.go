package main

import (
	"net/http"
    "./ctrl"
	"log"
	"html/template"
	"fmt"
    _ "github.com/go-sql-driver/mysql"
    "./web"
	)


func RegisterView(){
	//一次解析出全部模板
	tpl,err := template.ParseGlob("view/**/*")
	if nil!=err{
		log.Fatal(err)
	}
	//通过for循环做好映射
	for _,v := range tpl.Templates(){
		tplname := v.Name();
		fmt.Println("HandleFunc     "+v.Name())


		http.HandleFunc(tplname, func(w http.ResponseWriter,
			request *http.Request) {
			fmt.Println("parse     "+v.Name() + "==" + tplname)
			err := tpl.ExecuteTemplate(w,tplname,nil)
			if err!=nil{
				log.Fatal(err.Error())
			}
		})
	}

}

func main() {

	server := web.NewSdkHttpServer("im-server",web.MetricFilterBuilder)

	//绑定请求和处理函数
	_ = server.Route("POST", "/user/login", ctrl.UserLogin)//登录
	_ = server.Route("POST", "/user/register", ctrl.UserRegister)//注册
	_ = server.Route("GET", "/contact/loadcommunity", ctrl.LoadCommunity)//加载群聊
	_ = server.Route("POST", "/contact/loadfriend", ctrl.LoadFriend)//加载好友
	_ = server.Route("POST", "/contact/joincommunity", ctrl.JoinCommunity)//加入群聊
	_ = server.Route("POST", "/contact/createcommunity", ctrl.CreateCommunity)//创建群聊
	_ = server.Route("POST", "/contact/addfriend", ctrl.Addfriend)//添加好友
	_ = server.Route("GET", "/chat", ctrl.Chat)//聊天
	_ = server.Route("POST", "/attach/upload", ctrl.Upload)//上传
	
	//前端页面，注册模板
	tpl,err := template.ParseGlob("view/**/*")
	if nil!=err{
		log.Fatal(err)
	}
	//通过for循环做好映射
	for _,v := range tpl.Templates(){
		tplname := v.Name();
		fmt.Println("HandleFunc     "+v.Name())

		_ = server.Route("GET", tplname,func(c *web.Context){
				fmt.Println("parse     "+v.Name() + "==" + tplname)
				err := tpl.ExecuteTemplate(c.W,tplname,nil)
				if err!=nil{
					log.Fatal(err.Error())
				}
			})

	}

	//指定目录的静态文件
	http.Handle("/asset/",http.FileServer(http.Dir(".")))
	http.Handle("/mnt/",http.FileServer(http.Dir(".")))



	server.Start(":8080")
	
}