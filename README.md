# QRCode Service
 I'm a web app, using the go language.
 ![image](https://raw.githubusercontent.com/chuliam/QRCodeService/master/resource/demo.png)
 
# Pre condition
- vendor go tool
- Go environment

# How to use?
1. go get github.com/chuliam/QRCodeService
2. modify the conf/app.ini, enter your config info
3. go build
4. run

# About app.ini
[QRCODE]

AppName=QR.Code *Your application name*

AppVer=V1_release *Your application version*

AppPort=8080 *Your application port*

AppUri=http://localhost:8080 *Your application host & port*

AppLogLevel=0 *log level*
#0:debug 1:info 2:warn 3:error 4:panic 5:fatal 6:none

AppLogFileFormat=ByHour *log file name format*

#ByDay:"2006-01-02",ByHour:"2006-01-02-15",ByMonth:"2006-01"

IndexFileUri=./index.html *index file url*

QRCodeSaveDir=./public/ *save qrcode png data dir*
