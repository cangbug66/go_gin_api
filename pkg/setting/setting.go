package setting

import (
    "github.com/go-ini/ini"
    "log"
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

func Setup(){
    Cfg,err:=ini.Load("./conf/app.ini")
    if err!=nil{
        log.Fatalf("Fail to parse 'conf/app.ini:': %v",err)
    }

    //App
    err=Cfg.Section("app").MapTo(&AppSetting)
    if err != nil {
        log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
    }

    //server
    err = Cfg.Section("server").MapTo(ServerSetting)
    if err != nil {
        log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
    }

    //database
    err = Cfg.Section("database").MapTo(DatabaseSetting)
    if err != nil {
        log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
    }
}

