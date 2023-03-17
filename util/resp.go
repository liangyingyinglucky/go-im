package util

import (
	"net/http"
	"encoding/json"
	"log"
)

//返回数据结构体
type H struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
	Data interface{} `json:"data,omitempty"`//非必须
	Rows interface{} `json:"rows,omitempty"`
	Total interface{} `json:"total,omitempty"`
}
//错误输出
func RespFail(w http.ResponseWriter,msg string){
	Resp(w,-1,nil,msg)
}
//正确输出
func RespOk(w http.ResponseWriter,data interface{},msg string){
	Resp(w,0,data,msg)
}
//有数据分页
func RespOkList(w http.ResponseWriter,lists interface{},total interface{}){
	//分页数目,
	RespList(w,0,lists,total)
}

func Resp(w http.ResponseWriter,code int,data interface{},msg string)  {
	//定义输出类型
	w.Header().Set("Content-Type","application/json")
	//设置200状态
	w.WriteHeader(http.StatusOK)
	//输出
	h := H{
		Code:code,
		Msg:msg,
		Data:data,
	}
	ret,err := json.Marshal(h)
	if err!=nil{
		log.Println(err.Error())
	}
	//输出
	w.Write(ret)
}
func RespList(w http.ResponseWriter,code int,data interface{},total interface{})  {

	w.Header().Set("Content-Type","application/json")
	//设置200状态
	w.WriteHeader(http.StatusOK)
 	h := H{
		Code:code,
		Rows:data,
		Total:total,
	}
	ret,err := json.Marshal(h)
	if err!=nil{
		log.Println(err.Error())
	}
	//输出
	w.Write(ret)
}