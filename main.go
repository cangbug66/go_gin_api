package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "go_gin_api/app/routers"
    "go_gin_api/appInit"
    "log"
    "net/http"
)

func main() {
    defer appInit.CloseDB()
    gin.SetMode(appInit.ServerSetting.RunMode)
    routersInit:=routers.InitRouter()

    //readTimeout := setting.ServerSetting.ReadTimeout
    //writeTimeout := setting.ServerSetting.WriteTimeout
    addr := fmt.Sprintf(":%d", appInit.ServerSetting.HttpPort)
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
