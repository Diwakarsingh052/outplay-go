package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
)

var tokenStr = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhcGkgcHJvamVjdCIsInN1YiI6IjEyMzQ1Njc4OSIsImV4cCI6MTYzNTM0Njg4OSwiaWF0IjoxNjM1MzQzODg5LCJBdXRob3JpemVkIjpbIkFETUlOIl19.C2do8TnPGl0ZUKe4ZntHXukoFl23xJws0gtWU-5V0LjDaFwIaZuJkkXvbrttlZ9hQ7aeKVF_wYvcX3CskJHHL5KKrNNBTaBV8fj7lvH8EL01U1cK11UqkQ1gwRx7Ky1HJmj7ENJ5Hu8hAtoS3BWzTkI2kOpMiICrp7vpJgCJuHYc1P5sxQk1SBdrDlFrKPJTElg9hfYaQNfTzHUgHw_80gYf0RhReKjQIPwAQe_vYoK7NoBtQLr8uan9s-8-B_-I9zNL2N0vLrdXtVMSl-N_8CKJzo3GT62ItKFwxNH0Ux4tG4uOdmWVrFdw8ab525O4tsUyrKeYgPnbQFSmAb6kUg`
var privateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAvMAHb0IoLvoYuW2kA+LTmnk+hfnBq1eYIh4CT/rMPCxgtzjq
U0guQOMnLg69ydyA5uu37v6rbS1+stuBTEiMQl/bxAhgLkGrUhgpZ10Bt6GzSEgw
QNloZoGaxe4p20wMPpT4kcMKNHkQds3uONNcLxPUmfjbbH64g+seg28pbgQPwKFK
tF7bIsOBgz0g5Ptn5mrkdzqMPUSy9k9VCu+R42LH9c75JsRzz4FeN+VzwMAL6yQn
ZvOi7/zOgNyxeVia8XVKykrnhgcpiOn5oaLRBzQGN00Z7TuBRIfDJWU21qQN4Cq7
keZmMP4gqCVWjYneK4bzrG/+H2w9BJ2TsmMGvwIDAQABAoIBAFQmQKpHkmavNYql
6POaksBRwaA1YzSijr7XJizGIXvKRSwqgb2zdnuTSgpspAx09Dr/aDdy7rZ0DAJt
fk2mInINDottOIQm3txwzTS58GQQAT/+fxTKWJMqwPfxYFPWqbbU76T8kXYna0Gs
OcK36GdMrgIfQqQyMs0Na8MpMg1LmkAxuqnFCXS/NMyKl9jInaaTS+Kz+BSzUMGQ
zebfLFsf2N7sLZuimt9zlRG30JJTfBlB04xsYMo734usA2ITe8U0XqG6Og0qc6ev
6lsoM8hpvEUsQLcjQQ5up7xx3S2stZJ8o0X8GEX5qUMaomil8mZ7X5xOlEqf7p+v
lXQ46cECgYEA2lbZQON6l3ZV9PCn9j1rEGaXio3SrAdTyWK3D1HF+/lEjClhMkfC
XrECOZYj+fiI9n+YpSog+tTDF7FTLf7VP21d2gnhQN6KAXUnLIypzXxodcC6h+8M
ZGJh/EydLvC7nPNoaXx96bohxzS8hrOlOlkCbr+8gPYKf8qkbe7HyxECgYEA3U6e
x9g4FfTvI5MGrhp2BIzoRSn7HlNQzjJ71iMHmM2kBm7TsER8Co1PmPDrP8K/UyGU
Q25usTsPSrHtKQEV6EsWKaP/6p2Q82sDkT9bZlV+OjRvOfpdO5rP6Q95vUmMGWJ/
S6oimbXXL8p3gDafw3vC1PCAhoaxMnGyKuZwlM8CgYEAixT1sXr2dZMg8DV4mMfI
8pqXf+AVyhWkzsz+FVkeyAKiIrKdQp0peI5C/5HfevVRscvX3aY3efCcEfSYKt2A
07WEKkdO4LahrIoHGT7FT6snE5NgfwTMnQl6p2/aVLNun20CHuf5gTBbIf069odr
Af7/KLMkjfWs/HiGQ6zuQjECgYEAv+DIvlDz3+Wr6dYyNoXuyWc6g60wc0ydhQo0
YKeikJPLoWA53lyih6uZ1escrP23UOaOXCDFjJi+W28FR0YProZbwuLUoqDW6pZg
U3DxWDrL5L9NqKEwcNt7ZIDsdnfsJp5F7F6o/UiyOFd9YQb7YkxN0r5rUTg7Lpdx
eMyv0/UCgYEAhX9MPzmTO4+N8naGFof1o8YP97pZj0HkEvM0hTaeAQFKJiwX5ijQ
xumKGh//G0AYsjqP02ItzOm2mWnbI3FrNlKmGFvR6VxIZMOyXvpLofHucjJ5SWli
eYjPklKcXaMftt1FVO4n+EKj1k1+Tv14nytq/J5WN+r4FBlNEYj/6vg=
-----END RSA PRIVATE KEY-----
`)

func validate() {
	type claims struct {
		Username string `json:"username"`
		jwt.RegisteredClaims
		Roles []string `json:"roles"`
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		log.Fatalln(err)
	}
	_ = privateKey
	var c claims
	token, err := jwt.ParseWithClaims(tokenStr, &c, func(token *jwt.Token) (interface{}, error) {
		//return privateKey.Public(), nil
		return []byte("hello"), nil
	})

	if err != nil {
		fmt.Println("parsing token ", err)
		return
	}

	if !token.Valid {
		fmt.Println("invalid token")
		return
	}

	fmt.Println(c)

}

func main() {
	validate()
}
