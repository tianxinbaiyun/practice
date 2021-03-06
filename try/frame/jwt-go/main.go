package main

import (
	"fmt"

	"time"

	"github.com/dgrijalva/jwt-go"
)

// SignNameSceret
var (
	SignNameSceret = "aweQurt178BNI"
)

func main() {
	fmt.Println("Hello World!")

	tokenString, err := createJwt()
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}

	fmt.Println("tokenString:", tokenString)
	claims := parseJwt(tokenString)
	fmt.Println("claims:", claims)

}

//验证
//在调用Parse时，会进行加密验证，同时如果提供了exp，会进行过期验证；
//如果提供了iat，会进行发行时间验证;如果提供了nbf，会进行发行时间验证．

//创建 tokenString
func createJwt() (string, error) {
	//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//		"foo": "bar",
	//		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),

	//	})

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["foo"] = "bar"
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(SignNameSceret))
	return tokenString, err
}

//解析tokenString
func parseJwt(tokenString string) jwt.MapClaims {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(SignNameSceret), nil
	})

	var claims jwt.MapClaims
	var ok bool
	if claims, ok = token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("claims[\"foo\"]=", claims["foo"], "claims[\"nbf\"]=", claims["nbf"])
	} else {
		fmt.Println("err:", err)
	}

	return claims
}
