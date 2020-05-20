package jwt

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "go_gin_api/app/pkg/app"
    "go_gin_api/app/pkg/e"
    "go_gin_api/app/pkg/util"
    "net/http"
    "time"
)

func JWT() gin.HandlerFunc {
    return func(c *gin.Context){
        appG:=app.Gin{C:c}
        code := e.SUCCESS
        token:=c.DefaultQuery("token","")
        fmt.Println("this")
        if token == ""{
            code = e.INVALID_PARAMS
        }else {
            claims,err:=util.ParseToken(token)
            if err!=nil{
               code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
            }else if time.Now().Unix() > claims.ExpiresAt{
               code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
            }
        }
        if code != e.SUCCESS{
            appG.Response(http.StatusUnauthorized,code,gin.H{"token":"error"})
            c.Abort()
            return
        }
       c.Next()
    }
}