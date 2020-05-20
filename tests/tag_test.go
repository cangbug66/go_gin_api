package tests

import (
    "github.com/gavv/httpexpect/v2"
    "github.com/gin-gonic/gin"
    "go_gin_api/app/routers"
    "net/http"
    "testing"
)

var e *httpexpect.Expect
var engine *gin.Engine

func init() {
    //appInit.Cfg,_=ini.Load("../config/app.ini")
    gin.SetMode(gin.TestMode)
    // Create new gin instance
    engine = gin.New()
    routers.Route_v1(engine)
}
func WithConfig(t *testing.T)  {
    // Create httpepect instance
    e = httpexpect.WithConfig(httpexpect.Config{
        Client: &http.Client{
            Transport: httpexpect.NewBinder(engine),
            Jar:       httpexpect.NewJar(),
        },
        Reporter: httpexpect.NewAssertReporter(t),
        Printers: []httpexpect.Printer{
            httpexpect.NewDebugPrinter(t, true),
        },
    })
}


func TestGetTagDetail(t *testing.T) {
    WithConfig(t)
    e.GET("/api/v1/tags/detail").WithQuery("id", 48).
      Expect().
      Status(http.StatusOK).JSON().Object().ValueEqual("code", 200)
}
