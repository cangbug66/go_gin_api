package app

import (
    "fmt"
    "github.com/astaxie/beego/validation"
    "github.com/gin-gonic/gin"
    "go_gin_api/pkg/e"
    "net/http"
)

type Gin struct {
    C *gin.Context
}

type Response struct {
    Code int         `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data"`
}

func(g *Gin) Response(httpCode int,errCode int,data interface{}) {
    //g.C.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
    //g.C.Header("Access-Control-Allow-Origin", "*")
    g.C.JSON(httpCode,Response{Code:errCode,Msg:e.GetMsg(errCode),Data:data})
    return
}

func MarkErrors(errors []*validation.Error) {
    for _, err := range errors {
       fmt.Println(fmt.Sprintf("key:%v msg:%v |",err.Key, err.Message))
    }

    return
}

func(g *Gin) BindFormAndValid(form interface{}) (int, int) {
    if err:=g.C.ShouldBind(form);err!=nil{
        return http.StatusBadRequest, e.INVALID_PARAMS
    }

    valid:= validation.Validation{}
    check,err:=valid.Valid(form)
    if err != nil {
        fmt.Println("err:",err.Error())
        return http.StatusInternalServerError, e.ERROR
    }
    if !check {
        MarkErrors(valid.Errors)
        return http.StatusBadRequest, e.INVALID_PARAMS
    }

    return http.StatusOK, e.SUCCESS

}