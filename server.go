package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Room struct {
	Password string `json:"password"`
	PaticNum int    `json:"particNum"`
}

func main() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	e.GET("/myPage", hello) // ローカル環境の場合、http://localhost:1323/ にGETアクセスされるとhelloハンドラーを実行する

	//e.POST("/createRoom", createRoom)
	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))
}

// ハンドラーを定義
func hello(c echo.Context) error {

	var fukami string = "Kensuke"
	return c.String(http.StatusOK, fukami)
}

// func createRoom(c echo.Context) error {
// 	password := c.FormValue("password")
// 	particNum := c.FormValue("particNum")
// 	int_particNum, _ := strconv.Atoi(particNum)

// 	useDB.CreateRoom(password, int_particNum, 0, "step")
// 	return c.String(http.StatusOK, "OK")
// }
