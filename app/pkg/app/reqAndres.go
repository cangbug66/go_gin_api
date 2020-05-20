package app


import (
    "fmt"
    "github.com/astaxie/beego/validation"
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "github.com/go-playground/locales/zh"
    "github.com/go-playground/universal-translator"
    "github.com/pkg/errors"
    "go_gin_api/app/pkg/e"
    "go_gin_api/appInit"
    "gopkg.in/go-playground/validator.v9"
    zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
    "log"
    "net/http"
    "strings"
)
type Gin struct {
    C *gin.Context
}

func GetAppSerive(c *gin.Context) *Gin {
    AppG= &Gin{C:c}
    return  AppG
}

var AppG *Gin

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

func(g *Gin) SuccessResponse(data interface{})   {
    g.Response(http.StatusOK,e.SUCCESS,data)
}

func(g *Gin) ParamErrorResponse(msg interface{})  {
    g.Response(http.StatusBadRequest,e.INVALID_PARAMS,msg)
}
func(g *Gin) SystemErrorResponse(msg interface{})  {
    g.Response(http.StatusInternalServerError,e.ERROR,msg)
}

var trans ut.Translator

func init()  {
    validateInit()
}

func validateInit(){
    zhObj := zh.New()
    uni := ut.New(zhObj, zhObj)
    var ok bool
    if trans, ok = uni.GetTranslator("zh");!ok{
      log.Fatalf("validator GetTranslator not found")
    }
    validate:=binding.Validator.Engine().(*validator.Validate)
    err:=zh_translations.RegisterDefaultTranslations(validate, trans)
    if err!=nil{
        log.Fatalf("validator zh_translations err:%v",err)
    }
}

/**
    form 必须传指针
 */
func(g *Gin) Validate1(form interface{}) error {
    if errs:=g.C.ShouldBind(form);errs!=nil{
        if errs,ok := errs.(validator.ValidationErrors);ok{
           msg:=[]string{}
           for _, err := range errs {
               msg=append(msg,err.Translate(trans))
           }
           str:=strings.Join(msg,"，")
           return errors.New(str)
        }
    }
    return nil
}
func(g *Gin) Validate2(form interface{})  {
    if errs:=g.C.ShouldBind(form);errs!=nil{

    }
    //return
    zhObj := zh.New()
    uni := ut.New(zhObj, zhObj)
    trans, _ := uni.GetTranslator("zh")
    validate := validator.New()

    err:=zh_translations.RegisterDefaultTranslations(validate, trans)
    if err!=nil{
        g.ParamErrorResponse(err.Error())
    }

    err = validate.Struct(form)
    fmt.Println(11)
    fmt.Println(err)
    if err != nil {
        errs := err.(validator.ValidationErrors)
        msg:=[]string{}
        for _, err := range errs {
            msg=append(msg,err.Translate(trans))
        }
        str:=strings.Join(msg,"，")
        g.ParamErrorResponse(str)
    }
}



func MarkErrors(errors []*validation.Error) {
    for _, err := range errors {
       fmt.Println(fmt.Sprintf("key:%v msg:%v |",err.Key, err.Message))
    }

    return
}

func(g *Gin) BindFormAndValid(form interface{}) (int, int) {
    if err:=g.C.Bind(form);err!=nil{
        fmt.Println("bind err:"+err.Error())
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

func (g *Gin) GetOffset(page int) int {
    p:=0
    if page>0{
        p=(page-1) * g.GetPageSize()
    }
    return p
}
func (g *Gin) GetPageSize() int {
    return appInit.AppSetting.PageSize
}