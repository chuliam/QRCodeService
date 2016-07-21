package main

import (
	"github.com/lunny/tango"
	"github.com/lunny/log"
	"QRCodeService/Actions"
	"QRCodeService/setting"
)

func main() {
	setting.LoadConfig()
	var s log.ByType
	switch setting.AppLogFileFormat {
	case "ByHour":
		s=log.ByHour
	case "ByMonth":
		s=log.ByMonth
	default:
		s=log.ByDay
	}
	w := log.NewFileWriter(log.FileOptions{
		ByType:s,
		Dir:"./logs",
	})
	log.Std.SetOutput(w)
	l := log.New(w, "["+setting.AppName +"_"+ setting.AppVer+"]", log.Ldefault())
	l.SetOutputLevel(setting.AppLogLevel)
	t := tango.Classic(l)
	t.Any("/GetQR", new(Actions.GennerateQR))
	t.Get("*", tango.File(setting.IndexFileUri))
	t.Get("/QrCode/:name", tango.Dir(setting.QRCodeSaveDir))
	t.Run(setting.AppPort)
}
