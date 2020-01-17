package v1

import (
    "github.com/gin-gonic/gin"
    "go_gin_api/models"
    "go_gin_api/pkg/app"
    "go_gin_api/pkg/e"
    "net/http"
)

type AddTagForm struct {
    Name string `form:"name" valid:"Required;MaxSize(20)"`
    CreatedBy string `form:"create_by"`
    State int `form:"state"`
}

// @Summary Add article tag
// @tag.name  Tag标签
// @Accept  mpfd
// @Produce  json
// @Param name formData string true "Name"
// @Param state formData int false "State"
// @Param created_by formData string false "CreatedBy"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/add [post]
func AddTag(c *gin.Context){
    var form AddTagForm
    appG:=app.Gin{C:c}

    httCode,errCode:=appG.BindFormAndValid(&form)

    if errCode != e.SUCCESS{
       appG.Response(httCode,errCode,nil)
       return
    }

    id,err:=models.AddTag(form.Name,form.State,form.CreatedBy)
    if err!=nil{
        appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
        return
    }

    appG.Response(http.StatusOK, e.SUCCESS, id)
}
