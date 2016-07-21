package setting

import (
	"github.com/Unknwon/goconfig"
)

const (
	DEFAULT_SECTION = "QRCODE"
	AppConfPath = "./conf/app.ini"
)

var (
	Cfg     *goconfig.ConfigFile
)
var (
	AppVer string
	AppLogLevel int
	AppPort string
	AppUri string
	AppName string
	IndexFileUri string
	QRCodeSaveDir string
	AppLogFileFormat string
)

func LoadConfig() *goconfig.ConfigFile {
	var err error
	Cfg, err = goconfig.LoadConfigFile(AppConfPath)
	if err != nil {
		println("Fail to load configuration file: " + err.Error())
		//os.Exit(2)
	}
	AppName = Cfg.MustValue(DEFAULT_SECTION, "AppName")
	AppPort = ":" + Cfg.MustValue(DEFAULT_SECTION, "AppPort")
	AppVer = Cfg.MustValue(DEFAULT_SECTION, "AppVer")
	AppUri = Cfg.MustValue(DEFAULT_SECTION, "AppUri")
	AppLogFileFormat = Cfg.MustValue(DEFAULT_SECTION, "AppLogFileFormat")
	AppLogLevel = Cfg.MustInt(DEFAULT_SECTION, "AppLogLevel")
	IndexFileUri = Cfg.MustValue(DEFAULT_SECTION, "IndexFileUri")
	QRCodeSaveDir = Cfg.MustValue(DEFAULT_SECTION, "QRCodeSaveDir")
	return Cfg
}
