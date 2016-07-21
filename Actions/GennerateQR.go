package Actions

import (
	"github.com/lunny/tango"
	"github.com/satori/go.uuid"
	"github.com/skip2/go-qrcode"
	"time"
	"strings"
	"QRCodeService/setting"
)
type GennerateQR struct {
	tango.Json
	tango.Ctx
}

func (h *GennerateQR)Get()interface{}{
	m:=h.Form("msg")
	if m==""{
		return OutModel{Code:200,Msg:"没有内容，无法产生"}
	}
	name:=strings.Replace(uuid.NewV4().String(),"-","",-1)//替换guid中的“-”
	err := qrcode.WriteFile(m, qrcode.Medium, 256, "public/"+name)
	if err!=nil{
		println("出错了"+err.Error())
	}
	uri:=setting.AppUri+"/QrCode/"+name
	expTime:=time.Now()
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["url"]=uri
	m1["content"]=m
	m1["expDate"]=expTime.Format("2006-01-02")+" 23:59 PM"
	return OutModel{Code:200,Data:m1}
}

type OutModel struct  {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`

}

