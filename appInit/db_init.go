package appInit

import (
    "fmt"
    "github.com/jinzhu/gorm"
    "log"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
    ID int`gorm:"primary_key" json:"id"`
    CreatedOn  int `json:"created_on"`
    ModifiedOn int `json:"modified_on"`
    DeletedOn  int `json:"deleted_on"`
}

var Db *gorm.DB
func Db_Setup(){
    var err error
    Db, err = gorm.Open(DatabaseSetting.Type , fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
        DatabaseSetting.User,
        DatabaseSetting.Password,
        DatabaseSetting.Host,
        DatabaseSetting.DbName))
    if err != nil {
        log.Fatalf("db init err:%v",err)
    }
    Db.SingularTable(true)
    Db.DB().SetMaxIdleConns(10)
    Db.DB().SetMaxOpenConns(100)
    if ServerSetting.RunMode == "debug"{
        Db.LogMode(true)
    }
    gorm.DefaultTableNameHandler= func(db *gorm.DB, defaultTableName string) string {
        return DatabaseSetting.TablePrefix+defaultTableName
    }
}

func CloseDB()  {
    if err:=Db.Close();err!=nil{
        log.Fatalf("db close err:%v",err)
    }
}

func GetDb() *gorm.DB{
    return Db
}
