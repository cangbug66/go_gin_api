package v1

import (
    "github.com/gin-gonic/gin"
    "go_gin_api/app/forms"
    "go_gin_api/app/models"
    "go_gin_api/app/pkg/app"
    "go_gin_api/app/pkg/e"
    "go_gin_api/app/services"
    "net/http"
)





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
    var form forms.AddTagForm
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


// @Summary List article tag
// @tag.name  Tag标签
// @Produce  json
// @Param name query string true "Name"
// @Param email query string false "email"
// @Param page query int false "page"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags [get]
func GetTagList(c *gin.Context)  {
    appG:=app.GetAppSerive(c)
    var tagListForm forms.TagListForm1
    //var tagListForm TagListForm2
    if err:=appG.Validate1(&tagListForm);err!=nil{
        appG.ParamErrorResponse(err.Error())
        return
    }
    //appG.Validate2(&tagListForm)

    wheres:=gin.H{"name":tagListForm.Name}
    tagsList,err:=services.GetTagList(wheres,appG.GetOffset(tagListForm.Page),appG.GetPageSize())
    if err!=nil{
        appG.SystemErrorResponse(err.Error())
        return
    }
    appG.SuccessResponse(tagsList)
    return
}

func GetTagDetail(c *gin.Context)  {
    appG:=app.GetAppSerive(c)
    var tagDeatilForm forms.TagDeatilForm
    if err:=appG.Validate1(&tagDeatilForm);err!=nil{
        appG.ParamErrorResponse(err.Error())
        return
    }
    tagDetail,err:=services.GetTagDetail(tagDeatilForm.Id)
    if err!=nil{
        appG.SystemErrorResponse(err.Error())
        return
    }
    appG.SuccessResponse(tagDetail)
    return
}

