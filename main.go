package main

import (
	"net/http"
    "./ctrl"
	"log"
	"html/template"
	"fmt"
    _ "github.com/go-sql-driver/mysql"
	)

//注册模板
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
	//绑定请求和处理函数
	http.HandleFunc("/user/login", ctrl.UserLogin)//登录
	http.HandleFunc("/user/register", ctrl.UserRegister)//注册
	http.HandleFunc("/contact/loadcommunity", ctrl.LoadCommunity)//加载群聊
	http.HandleFunc("/contact/loadfriend", ctrl.LoadFriend)//加载好友
	http.HandleFunc("/contact/joincommunity", ctrl.JoinCommunity)//加入群聊
	http.HandleFunc("/contact/createcommunity", ctrl.CreateCommunity)//创建群聊
	http.HandleFunc("/contact/addfriend", ctrl.Addfriend)//添加好友
	http.HandleFunc("/chat", ctrl.Chat)//聊天
	http.HandleFunc("/attach/upload", ctrl.Upload)//上传
	//指定目录的静态文件
	http.Handle("/asset/",http.FileServer(http.Dir(".")))
	http.Handle("/mnt/",http.FileServer(http.Dir(".")))

	//前端页面
	RegisterView()

	http.ListenAndServe(":8080",nil)
}