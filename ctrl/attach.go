package ctrl

import (
	"net/http"
	"../util"
	"os"
	"strings"
	"fmt"
	"time"
	"math/rand"
	"io"
	"../web"
	//"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

const (
AccessKeyId=""
AccessKeySecret=""
EndPoint=""
Bucket="test"
)

//创建上传文件地址
func init(){
	os.MkdirAll("./mnt",os.ModePerm)
}
func Upload(c *web.Context){
	UploadLocal(c.W,c.R)
}

//1.存储位置 ./mnt,需要确保已经创建好
//2.url格式 /mnt/xxxx.png  需要确保网络能访问/mnt/
func UploadLocal(writer http.ResponseWriter,
	request * http.Request){
	//上传文件
    srcfile,head,err:=request.FormFile("file")
    if err!=nil{
    	util.RespFail(writer,err.Error())
	}

	suffix := ".png"
	//文件名称包含后缀 xx.xx.png
	ofilename := head.Filename
	tmp := strings.Split(ofilename,".")
	if len(tmp)>1{
		suffix = "."+tmp[len(tmp)-1]
	}
	//前端指定filetype
	filetype := request.FormValue("filetype")
	if len(filetype)>0{
		suffix = filetype
	}
	//保存的文件名
    filename := fmt.Sprintf("%d%04d%s",
    	time.Now().Unix(), rand.Int31(),
    	suffix)



    //本地文件创建--------
    dstfile,err:= os.Create("./mnt/"+filename)
    if err!=nil{
    	util.RespFail(writer,err.Error())
    	return
	}

	//todo 将源文件内容copy到新文件
	_,err = io.Copy(dstfile,srcfile)
	if err!=nil{
		util.RespFail(writer,err.Error())
		return
	}
	//todo 将新文件路径转换成url地址

	url := "/mnt/"+filename
	//todo 响应到前端
	util.RespOk(writer,url,"")


	//oss文件创建------------
	//初始化ossclient
	/*client,err:=oss.New(EndPoint,AccessKeyId,AccessKeySecret)
	if err!=nil{
		util.RespFail(writer,err.Error())
		return
	}
	//获得bucket
	bucket,err := client.Bucket(Bucket)
	if err!=nil{
		util.RespFail(writer,err.Error())
		return
	}
	//通过bucket上传
	err=bucket.PutObject(filename,srcfile)
	if err!=nil{
		util.RespFail(writer,err.Error())
		return
	}
	//获得url地址
	url := "http://"+Bucket+"."+EndPoint+"/"+filename
	util.RespOk(writer,url,"")
*/

	}


