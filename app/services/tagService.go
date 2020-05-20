package services

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "go_gin_api/app/models"
    "go_gin_api/appInit"

)

func GetTagList(wheres map[string]interface{},offset int,pagesize int ) (interface{},error) {
    var tags models.TagList
    fmt.Println(offset)
    fmt.Println(pagesize)
    if err:=appInit.GetDb().Where(wheres).Offset(offset).Limit(pagesize).Order("id desc").Find(&tags).Error;err!=nil{
        if err != nil && err != gorm.ErrRecordNotFound {
            return nil, err
        }
    }
    var count int
    if err:=appInit.Db.Model(&models.Tag{}).Where(wheres).Count(&count).Error;err!=nil{
        if err != nil && err != gorm.ErrRecordNotFound {
            return nil, err
        }
    }
    return gin.H{"list":tags,"count":count},nil
}

func GetTagDetail(id interface{}) (interface{},error){
    var tagDetail models.Tag
    if err:=appInit.Db.Where("id=?",id).Find(&tagDetail).Error;err!=nil{
        if err != nil && err != gorm.ErrRecordNotFound {
            return nil, err
        }
    }
    if tagDetail.ID>0{
        return  tagDetail,nil
    }
    return gin.H{},nil
}
