package Diy_validator

import (
    "fmt"
    "gopkg.in/go-playground/validator.v9"
)


func Integer(fl validator.FieldLevel) bool {
    fmt.Println("gg")
    fmt.Println(fl.Field().Int())
    return fl.Field().String() == "dd"
}
