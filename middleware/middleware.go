package middleware

import (
	"fmt"
	"jiuban/db"
	"time"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

var mySecret = []byte("secret")

var J = jwt.New(jwt.Config{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	},
	Expiration: true,
	// Extract by the "token" url.
	// There are plenty of options.
	// The default jwt's behavior to extract a token value is by
	// the `Authorization: Bearer $TOKEN` header.
	Extractor: jwt.FromAuthHeader,
	// When set, the middleware verifies that tokens are
	// signed with the specific signing algorithm
	// If the signing method is not constant the `jwt.Config.ValidationKeyGetter` callback
	// can be used to implement additional checks
	// Important to avoid security issues described here:
	// https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
	SigningMethod: jwt.SigningMethodHS256,
})

var c = jwt.New(jwt.Config{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	},
	Expiration: true,
	// Extract by the "token" url.
	// There are plenty of options.
	// The default jwt's behavior to extract a token value is by
	// the `Authorization: Bearer $TOKEN` header.
	Extractor: jwt.FromParameter("token"),
	// When set, the middleware verifies that tokens are
	// signed with the specific signing algorithm
	// If the signing method is not constant the `jwt.Config.ValidationKeyGetter` callback
	// can be used to implement additional checks
	// Important to avoid security issues described here:
	// https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
	SigningMethod: jwt.SigningMethodHS256,
})

// generate token to use.
func GetTokenHandler(ctx iris.Context, email string) string {
	var C = db.NewClient()

	now := time.Now()
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   now.Add(15 * time.Minute).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(mySecret)

	intcmd := C.Exists(email)
	if intcmd != nil {
		intcmd := C.Del(email)
		fmt.Println(intcmd.Result())
		err := C.Set(email, tokenString, 15*time.Minute).Err() //SET key value 0 數字代表過期秒數，在這裡0為永不過期
		if err != nil {
			panic(err)
		}
	} else {
		err := C.Set(email, tokenString, 15*time.Minute).Err() //SET key value 0 數字代表過期秒數，在這裡0為永不過期
		if err != nil {
			panic(err)
		}
	}

	val, err := C.Get(email).Result() //GET key value
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
	return tokenString

}

func AuthToken(ctx iris.Context) string {
	u := ctx.Values().Get("jwt").(*jwt.Token) //获取 token 信息
	email, _ := u.Claims.(jwt.MapClaims)["email"].(string)
	return email
}

func DeleteToken(ctx iris.Context, email string) {
	var C = db.NewClient()

	intcmd := C.Exists(email)
	if intcmd != nil {
		C.Del(email)
	}
}
