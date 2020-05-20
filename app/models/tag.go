package models

import (
    "github.com/jinzhu/gorm"
    "go_gin_api/appInit"
)

type Tag struct {
    appInit.Model

    Name string `json:"name"`
    CreatedBy string `json:"created_by"`
    ModifiedBy string `json:"modified_by"`
    State int `json:"state"`
}

type TagList []Tag

func ExistTagByName(name string) (bool,error){
    var tag Tag
    err:=appInit.GetDb().Select("id").Where("name=?",name).First(&tag).Error
    if err!=nil &&err!=gorm.ErrRecordNotFound{
        return false,err
    }
    if tag.ID > 0{
        return true,nil
    }

    return false,nil
}

func AddTag(name string,state int,createdBy string) (int,error){
    tag := Tag{
        Name:      name,
        State:     state,
        CreatedBy: createdBy,
    }
    if err := appInit.GetDb().Create(&tag).Error; err != nil {
        return 0,err
    }
    return tag.ID,nil
}

