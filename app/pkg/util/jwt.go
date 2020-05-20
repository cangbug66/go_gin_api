package util

import (
    "github.com/dgrijalva/jwt-go"
    "go_gin_api/appInit"
    "time"
)

type MyCustomClaims struct {
    Username string `json:"username"`
    Password  string `json:"password "`
    jwt.StandardClaims
}

func GenerateToken(username string,password string) (string,error){

    nowTime:=time.Now()
    expireTime:=nowTime.Add(3 * time.Hour).Unix()
    claims:=MyCustomClaims{
       username,
      password,
       jwt.StandardClaims {
           ExpiresAt : expireTime,
           Issuer : "abc",
       },
    }



    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    jwtSecret:=appInit.AppSetting.JwtSecret
    ss, err := token.SignedString([]byte(jwtSecret))
    return ss, err
}

func ParseToken(tokenString string) (*MyCustomClaims,error){
    tokenClaims, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(appInit.AppSetting.JwtSecret), nil
    })
    if err!=nil{
        //fmt.Println(err.Error())
        //if err ,ok := err.(*jwt.ValidationError);ok {
        //   if err.Errors & jwt.ValidationErrorMalformed != 0 {
        //       return nil,err
        //   }
        //   if err.Errors & (jwt.ValidationErrorExpired | jwt.ValidationErrorNotValidYet) != 0 {
        //       return nil,err
        //   }
        //}
        return nil,err
    }
    if claims, ok := tokenClaims.Claims.(*MyCustomClaims); ok && tokenClaims.Valid {
        //ExpiresAt:=time.Unix(claims.StandardClaims.ExpiresAt,0).Format("2006-01-02 15:04:05")
        //fmt.Printf("姓名：%v 时间： %v", claims.Username, ExpiresAt)
        return claims,nil
    }
    return nil,err
}