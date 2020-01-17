package test

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "go_gin_api/pkg/util"
)

func TestItem(c *gin.Context){
    //GenerateToken(c)
    ParseToken(c)
}

func GenerateToken(c *gin.Context){
    totken,err:=util.GenerateToken("lisis","123")
    if err!=nil{
        fmt.Println(err)
    }
    fmt.Println(totken)

    c.JSON(200,gin.H{"d":"ok"})
}

func ParseToken(c *gin.Context){
    token:=c.Query("token")
   clam,err:= util.ParseToken(token)
    //util.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Imxpc2lzIiwicGFzc3dvcmQgIjoiMTIzIiwiZXhwIjoxNTc5MjU5MTAwLCJpc3MiOiJhYmMifQ.Ovo4JrVwPWRP1d7v7rVrSK0nOAkGRb6aWKZOb8pKfJM")
    if err!=nil{
       //fmt.Println(err)
    }
    fmt.Println(clam)

    c.JSON(200,gin.H{"token":token})
}

