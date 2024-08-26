package utils
import (
    "github.com/dgrijalva/jwt-go"
    "time"
)
var jwtSecret=[]byte("your_secret_key")
func GenerateJWT(username string) (string, error){
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username":username,
        "exp":time.Now().Add(time.Hour*24).Unix(),
    })
    tokenString, err := token.SignedString(jwtSecret)
    if err!= nil{
    return "", err
    }
    return tokenString, nil
}
