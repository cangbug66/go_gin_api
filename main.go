package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "go_gin_api/models"
    "go_gin_api/pkg/setting"
    "go_gin_api/routers"
    "log"
    "net/http"
)

func init(){
    setting.Setup()
    models.Setup()

}

func main() {
    defer models.CloseDB()
    gin.SetMode(setting.ServerSetting.RunMode)
    routersInit:=routers.InitRouter()

    //readTimeout := setting.ServerSetting.ReadTimeout
    //writeTimeout := setting.ServerSetting.WriteTimeout
    addr := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
    //maxHeaderBytes := 1 << 20
    server:=&http.Server{
        Addr:addr,
        Handler:routersInit,
        //ReadTimeout:    readTimeout,
        //WriteTimeout:   writeTimeout,
        //MaxHeaderBytes: maxHeaderBytes,
    }
    err:=server.ListenAndServe()
    if err!=nil {
        log.Fatal(fmt.Sprintf("server start err: %s",err))
    }
}
