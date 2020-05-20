package appInit

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/go-ini/ini"
    "log"
    "os"
    "time"
)

type App struct {
    PageSize int
    JwtSecret string
}
var AppSetting=&App{}

type Server struct {
    RunMode string
    HttpPort int
    ReadTimeout time.Duration
    WriteTimeout time.Duration
}
var ServerSetting=&Server{}

type Database struct {
    Type string
    User string
    Password string
    Host string
    DbName string
    TablePrefix string
}
var DatabaseSetting=&Database{}
var Cfg *ini.File

func Config_Setup(){
    var err error
    var path string
    if gin.Mode() == "test" {
        path="../config/app.ini"
    }else {
        path = "./config/app.ini"
    }
    fmt.Println(os.Getwd())
    Cfg,err=ini.Load(path)
    if err!=nil{
        log.Fatalf("Fail to parse 'config/app.ini:': %v",err)
    }

    //App
    MapTo("app",AppSetting)
    //server
    MapTo("server",ServerSetting)
    //database
    MapTo("database",DatabaseSetting)
}

func MapTo(section string,setting interface{}){
    err := Cfg.Section(section).MapTo(setting)
    if err != nil {
        log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
    }
}



