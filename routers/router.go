package routers

import (
    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
    _ "go_gin_api/docs"
    "go_gin_api/middleware/jwt"
    "go_gin_api/routers/api/test"
    "go_gin_api/routers/api/v1"
    "net/http"
)

func InitRouter() *gin.Engine{
    r:=gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    route(r)

    return r
}

func route(r *gin.Engine){

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    ApiV1:=r.Group("/api/v1")
    {
        ApiV1.GET("tags", func(c *gin.Context) {
            c.JSON(http.StatusOK,gin.H{"data":"success"})
        })
        ApiV1.POST("tags/add", v1.AddTag)
    }
    ApiAithV1:=r.Group("/api/auth/v1")
    ApiAithV1.Use(jwt.JWT())
    {
        ApiAithV1.GET("auth", func(c *gin.Context) {
            c.JSON(http.StatusOK,gin.H{"data":"success"})
        })
    }
    r.GET("test",test.TestItem)
}
