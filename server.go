package main

import (
	"fmt"
	useDB "gamename-back-end/src/use_DB"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Room struct { //TODO　create_dbと被るからそこを考えよう
	WolfMode bool `json:"wolfMode"`
}

type Player struct {
	RoomId     string `json:"roomId"`
	PlayerName string `json:"playerName"`
	PlayerIcon int    `json:"playerIcon"`
}

type Theme struct {
	PlayerId string `json:"playerId"`
	RoomId   string `json:"roomId"`
	Text     string `json:"theme"`
}

type Hint struct {
	PlayerId string `json:"playerId"`
	RoomId   string `json:"roomId"`
	Hint     string `json:"hint"`
}

type DeleteHint struct { //TODO structの名前と型の修正
	RoomId string   `json:"roomId"`
	Hint   []string `json:"hint"`
}
type DecideTheme struct {
	RoomId           string `json:"roomId"`
	HowToDecideTheme int    `json:"howToDecideTheme"`
}
type Game struct {
	RoomId string `json:"roomId"`
}
type Answer struct {
	RoomId string `json:"roomId"`
	Answer string `json:"answer"`
}
type IsCorrect struct {
	RoomId    string `json:"roomId"`
	IsCorrect bool   `json:"isCorrect"`
}
type Vote struct {
	PlayerId string `json:"playerId"`
	RoomId   string `json:"roomId"`
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// ローカル環境の場合、http://localhost:1323/
	e.GET("/is-room-exit", isRoomExit)
	e.GET("/partic-list", func(c echo.Context) error { //TODO関数の管理ときに修正
		playerList := getParticList(c)
		return c.JSON(http.StatusOK, playerList)
	})
	e.GET("/theme", getTheme)
	e.GET("/hint-list", func(c echo.Context) error {
		hintList := getHintList(c)
		return c.JSON(http.StatusOK, hintList)
	})
	e.GET("/step", getStep)
	e.GET("/get-role", getRole)
	e.GET("/answer", getAnswer)
	e.GET("/judgement-answer", getJudgement)
	e.POST("/create-room", createRoom)
	e.POST("/add-player", postAddPlayer)
	e.POST("/create-theme", postCreateTheme)
	e.POST("/create-hint", postCreateHint)
	e.POST("/delete-hint", postDeleteHint)
	e.POST("/start-game", postStartGame)
	e.POST("/update-answer", postUpdateAnswer)
	e.POST("/is-correct", postIsCorrect)
	e.POST("/initialize", postEndGame)
	e.POST("/how-decide-theme", postDecideTheme)
	e.POST("/vote", postVote)
	e.Logger.Fatal(e.Start(":1323"))
}

func isRoomExit(c echo.Context) error {
	roomId := c.QueryParam("roomId")
	exit := useDB.IsRoomExit(roomId)

	return c.JSON(http.StatusOK, exit)
}

func getParticList(c echo.Context) []useDB.PlayerInfo {
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
func getRole(c echo.Context) error {
	playerId := c.QueryParam("playerId")
	return c.JSON(http.StatusOK, useDB.GetRole(playerId))
}

func getAnswer(c echo.Context) error {
	roomId := c.QueryParam("roomId")
	answer := useDB.GetAnswer(roomId)
	return c.JSON(http.StatusOK, answer)
}
func getJudgement(c echo.Context) error {
	roomId := c.QueryParam("roomId")
	answer := useDB.JudgementAnswer(roomId)
	return c.JSON(http.StatusOK, answer)
}
func createRoom(c echo.Context) error {
	reqBody := new(Room)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	wolfMode := reqBody.WolfMode
	return c.String(http.StatusOK, useDB.CreateRoom(0, "theme", 0, 0, wolfMode, false, 0))
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
	roomId := reqBody.RoomId
	return c.JSON(http.StatusOK, useDB.CreateTheme(theme, playerId, roomId))
}
func postCreateHint(c echo.Context) error {
	reqBody := new(Hint)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	playerId := reqBody.PlayerId
	hint := reqBody.Hint
	roomId := reqBody.RoomId
	return c.JSON(http.StatusOK, useDB.CreateHint(hint, playerId, roomId))
}
func postDeleteHint(c echo.Context) error {
	reqBody := new(DeleteHint)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	hintList := reqBody.Hint
	roomId := reqBody.RoomId
	return c.JSON(http.StatusOK, useDB.DeleteHint(hintList, roomId))
}
func postStartGame(c echo.Context) error {
	reqBody := new(Game)
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
func postUpdateAnswer(c echo.Context) error {
	reqBody := new(Answer)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	roomId := reqBody.RoomId
	answer := reqBody.Answer

	return c.JSON(http.StatusOK, useDB.UpdateAnswer(answer, roomId))
}
func postIsCorrect(c echo.Context) error {
	reqBody := new(IsCorrect)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	roomId := reqBody.RoomId
	isCorrect := reqBody.IsCorrect

	return c.JSON(http.StatusOK, useDB.IsCorrect(roomId, isCorrect))
}
func postEndGame(c echo.Context) error {
	reqBody := new(Game)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	roomId := reqBody.RoomId

	return c.JSON(http.StatusOK, useDB.EndGame(roomId))
}

func postVote(c echo.Context) error {
	reqBody := new(Vote)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	playerId := reqBody.PlayerId
	roomId := reqBody.RoomId
	return c.JSON(http.StatusOK, useDB.Vote(playerId, roomId))
}

// $body = @{
//     password = "yourpass"
// } | ConvertTo-Json
// Invoke-RestMethod -Method POST -Uri http://localhost:1323/create-room -Body $body -ContentType "application/json"
//curl -d "roomId = cbBipgOwuA8wxu5XAXFW" -d "playerName = testman" -d "playerIcon = 3" http://localhost:1323/addPlayer
