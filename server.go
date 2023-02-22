package main

import (
	"fmt"
	useDB "gamename-back-end/src/use_DB"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Room struct { //TODO　create_dbと被るからそこを考えよう
	Password string `json:"password"`
}

type Player struct {
	RoomId     string `json:"roomId"`
	PlayerName string `json:"playerName"`
	PlayerIcon int    `json:"playerIcon"`
}

type Theme struct {
	PlayerId string `json:"playerId"`
	Text     string `json:"theme"`
}

type Hint struct {
	PlayerId string `json:"playerId"`
	Hint     string `json:"hint"`
}

type DeleteHint struct { //TODO structの名前と型の修正
	Hint []string `json:"hint"`
}
type DecideTheme struct {
	RoomId           string `json:"roomId"`
	HowToDecideTheme int    `json:"howToDecideTheme"`
}
type StartGame struct {
	RoomId string `json:"roomId"`
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// ローカル環境の場合、http://localhost:1323/
	e.GET("/", hello)
	e.GET("/is-room-exit", isRoomExit)
	e.GET("/is-room-exit", isRoomExit)
	e.GET("/partic-list", func(c echo.Context) error { //TODO関数の管理ときに修正
		playerList := getParticList(c)
		return c.JSON(http.StatusOK, playerList)
	})
	e.GET("/theme:description", getTheme)
	e.GET("/hint-list", func(c echo.Context) error {
		hintList := getHintList(c)
		return c.JSON(http.StatusOK, hintList)
	})
	e.GET("/step", getStep)
	e.GET("/random-theme", getRandomTheme)
	e.GET("/get-role", getRole)
	e.POST("/create-room", createRoom)
	e.POST("/add-player", postAddPlayer)
	e.POST("/create-theme", postCreateTheme)
	e.POST("/create-hint", postCreateHint)
	e.POST("/delete-hint", postDeleteHint)
	e.POST("/start-game", postStartGame)
	e.Logger.Fatal(e.Start(":1323"))
}
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}
func isRoomExit(c echo.Context) error {
	roomId := c.QueryParam("roomId")
	exit := useDB.IsRoomExit(roomId)

	return c.JSON(http.StatusOK, exit)
}

func getParticList(c echo.Context) []useDB.PlayerNNNIcon {
	roomId := c.QueryParam("roomId")
	playerList := useDB.PlayerList(roomId)
	return playerList
}

func getTheme(c echo.Context) error {
	roomId := c.QueryParam("roomId")
	theme := useDB.GetTheme(roomId)
	return c.JSON(http.StatusOK, theme)
}

func getHintList(c echo.Context) []useDB.HintKey {
	roomId := c.QueryParam("roomId")
	return useDB.HintList(roomId)
}
func getStep(c echo.Context) error {
	roomId := c.QueryParam("roomId")
	fmt.Println(roomId)
	return c.JSON(http.StatusOK, useDB.GetStep(roomId))
}
func getRandomTheme(c echo.Context) error {
	return c.JSON(http.StatusOK, useDB.GetRandomTheme())
}
func getRole(c echo.Context) error {
	playerId := c.QueryParam("playerId")
	return c.JSON(http.StatusOK, useDB.GetRole(playerId))
}
func createRoom(c echo.Context) error {
	reqBody := new(Room)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	password := reqBody.Password
	return c.String(http.StatusOK, useDB.CreateRoom(password, 1, "theme", 0, 0))
}

func postAddPlayer(c echo.Context) error {
	reqBody := new(Player)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	roomId := reqBody.RoomId
	playerName := reqBody.PlayerName
	playerIcon := reqBody.PlayerIcon
	playerId := useDB.AddPlayer(roomId, playerName, playerIcon)
	return c.JSON(http.StatusOK, playerId)
}

func postCreateTheme(c echo.Context) error {
	reqBody := new(Theme)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	playerId := reqBody.PlayerId
	theme := reqBody.Text

	return c.JSON(http.StatusOK, useDB.CreateTheme(theme, playerId))
}
func postCreateHint(c echo.Context) error {
	reqBody := new(Hint)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	playerId := reqBody.PlayerId
	hint := reqBody.Hint
	return c.JSON(http.StatusOK, useDB.CreateHint(hint, playerId))
}
func postDeleteHint(c echo.Context) error {
	reqBody := new(DeleteHint)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	hintList := reqBody.Hint
	return c.JSON(http.StatusOK, useDB.DeleteHint(hintList))
}
func postStartGame(c echo.Context) error {
	reqBody := new(StartGame)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	roomId := reqBody.RoomId

	return c.JSON(http.StatusOK, useDB.StartGame(roomId))
}
func postDecideTheme(c echo.Context) error {
	reqBody := new(DecideTheme)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	roomId := reqBody.RoomId
	howToDecideTheme := reqBody.HowToDecideTheme
	return c.JSON(http.StatusOK, useDB.DecideTheme(roomId, howToDecideTheme))
}

// $body = @{
//     password = "yourpass"
//     particNum = 3
// } | ConvertTo-Json
// Invoke-RestMethod -Method POST -Uri http://localhost:1323/create-room -Body $body -ContentType "application/json"
//curl -d "roomId = cbBipgOwuA8wxu5XAXFW" -d "playerName = testman" -d "playerIcon = 3" http://localhost:1323/addPlayer
