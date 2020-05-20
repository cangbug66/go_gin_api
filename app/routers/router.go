package routers

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
    "go_gin_api/app/Diy_validator"
    "go_gin_api/appInit"
    _ "go_gin_api/docs"
    "gopkg.in/go-playground/validator.v9"
    "log"
)

func InitRouter() *gin.Engine{
    r:=gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    Route(r)

    return r
}

func Route(r *gin.Engine) *gin.Engine{


    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        if err:=v.RegisterValidation("integer", Diy_validator.Integer);err!=nil{
            log.Fatalf("validator integer err:%v",err)
        }
    }
    Route_v1(r)
    if appInit.ServerSetting.RunMode == "debug"{
        r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    }

    return r
}


func GinHandler(r *gin.Engine) *gin.Engine {
    // Add route to the gin engine
    r.GET("/example", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    // return gin engine with newly added route
    return r
}