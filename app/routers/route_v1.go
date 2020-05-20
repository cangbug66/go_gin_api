package routers

import (
    "github.com/gin-gonic/gin"
    "go_gin_api/app/controllers/v1"
    "go_gin_api/app/controllers/v1/test"
    "go_gin_api/app/middlewares/jwt"

    "net/http"
)

func Route_v1(r *gin.Engine){
    ApiV1:=r.Group("/api/v1")
    ApiV1.POST("tags/add",v1.AddTag )
    ApiV1.GET("tags",v1.GetTagList)
    ApiV1.GET("/tags/detail",v1.GetTagDetail)

    ApiAithV1:=r.Group("/api/auth/v1")
    ApiAithV1.Use(jwt.JWT())
    {
        ApiAithV1.GET("auth", func(c *gin.Context) {
            c.JSON(http.StatusOK,gin.H{"data":"success"})
        })
    }
    r.GET("test",test.TestItem)
}