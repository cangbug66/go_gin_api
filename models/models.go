package models

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "go_gin_api/pkg/setting"
    "log"
)

type Model struct {
    ID int`gorm:"primary_key" json:"id"`
    CreatedOn  int `json:"created_on"`
    ModifiedOn int `json:"modified_on"`
    DeletedOn  int `json:"deleted_on"`
}

var Db *gorm.DB
func Setup(){

    var err error
    Db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
        setting.DatabaseSetting.User,
        setting.DatabaseSetting.Password,
        setting.DatabaseSetting.Host,
        setting.DatabaseSetting.DbName))
    if err != nil {
        log.Println(err)
    }
    Db.SingularTable(true)
    Db.DB().SetMaxIdleConns(10)
    Db.DB().SetMaxOpenConns(100)

    gorm.DefaultTableNameHandler= func(db *gorm.DB, defaultTableName string) string {
        return setting.DatabaseSetting.TablePrefix+defaultTableName
    }

}

func CloseDB()  {
    defer Db.Close()
}

func GetDb() *gorm.DB{
    return Db
}