package main

import (
	"fmt"
	connectDB "gamename-back-end/pkg/connect_db"
	createDB "gamename-back-end/pkg/cruds/create"
	readDB "gamename-back-end/pkg/cruds/read"
	types "gamename-back-end/pkg/types"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// ローカル環境の場合、http://localhost:1323/
	e.GET("/is-mode-wolf", isModeWolf)
	e.GET("/is-room-exit", isRoomExit)
	e.GET("/partic-list", getParticList)
	e.GET("/partic-list-wolf", getParticListWolf)
	e.GET("/partic-list-vote", getVotePlayerList)
	e.GET("/theme", getTheme)
	e.GET("/hint-list", getHintList)
	e.GET("/step", getStep)
	e.GET("/get-role", getRole)
	e.GET("/get-role-wolf", getRoleWolf)
	e.GET("/answer", getAnswer)
	e.GET("/judgement-answer", getJudgement)
	e.GET("/vanish-wolf", getChoiceWolf)
	e.GET("/get-wolf-name", getWolfName)
	e.GET("/point", getPoint)
	e.GET("/result", getResult)

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
	e.POST("/judgement-wolf", postJudgementWolf)
	e.POST("/add-step", postAddStep)
	e.Logger.Fatal(e.Start(getPort()))
}
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":1323"
	} else {
		port = ":" + port
	}

	fmt.Println(port)
	return port
}
func isModeWolf(c echo.Context) error {
	roomId := c.QueryParam("roomId")
	return c.JSON(http.StatusOK, readDB.IsModeWolf(roomId))
}
func isRoomExit(c echo.Context) error {
	roomId := c.QueryParam("roomId")
	exit := readDB.IsRoomExit(roomId)

	return c.JSON(http.StatusOK, exit)
}
func getParticList(c echo.Context) error {
	roomId := c.QueryParam("roomId")
	playerList := readDB.PlayerList(roomId)
	return c.JSON(http.StatusOK, playerList)
}

func getParticListWolf(c echo.Context) error {
	roomId := c.QueryParam("roomId")
	playerListWolf := readDB.PlayerListWolf(roomId)
	return c.JSON(http.StatusOK, playerListWolf)
}

func getTheme(c echo.Context) error {
	roomId := c.QueryParam("roomId")
	theme := readDB.GetTheme(roomId)
	return c.JSON(http.StatusOK, theme)
}

func getHintList(c echo.Context) error {
	roomId := c.QueryParam("roomId")
	return c.JSON(http.StatusOK, readDB.HintList(roomId))
}
func getStep(c echo.Context) error {
	roomId := c.QueryParam("roomId")
	fmt.Println(roomId)
	return c.JSON(http.StatusOK, readDB.GetStep(roomId))
}
func getRole(c echo.Context) error {
	playerId := c.QueryParam("playerId")
	return c.JSON(http.StatusOK, readDB.GetRole(playerId))
}
func getRoleWolf(c echo.Context) error {
	playerId := c.QueryParam("playerId")
	return c.JSON(http.StatusOK, readDB.GetRoleWolf(playerId))
}
func getAnswer(c echo.Context) error {
	roomId := c.QueryParam("roomId")
	answer := readDB.GetAnswer(roomId)
	return c.JSON(http.StatusOK, answer)
}
func getJudgement(c echo.Context) error {
	roomId := c.QueryParam("roomId")
	answer := readDB.JudgementAnswer(roomId)
	return c.JSON(http.StatusOK, answer)
}
func getChoiceWolf(c echo.Context) error {
	roomId := c.QueryParam("roomId")
	return c.JSON(http.StatusOK, readDB.ChoiceWolf(roomId))
}
func getPoint(c echo.Context) error {
	roomId := c.QueryParam("roomId")
	fmt.Println(roomId)
	return c.JSON(http.StatusOK, readDB.PointCal(roomId))
}
func getVotePlayerList(c echo.Context) error {
	roomId := c.QueryParam("roomId")
	return c.JSON(http.StatusOK, readDB.VotePlayerList(roomId))
}

func getWolfName(c echo.Context) error {
	roomId := c.QueryParam("roomId")
	return c.JSON(http.StatusOK, readDB.WolfName(roomId))
}
func getResult(c echo.Context) error {
	roomId := c.QueryParam("roomId")
	return c.JSON(http.StatusOK, readDB.GetResult(roomId))
}
func createRoom(c echo.Context) error {
	reqBody := new(types.Room)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	wolfMode := reqBody.WolfMode

	ctx, client, err := connectDB.ConnectDB()
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}

	var id = createDB.CreateRoom(ctx, client, 0, "theme", 0, 0, wolfMode, false, 0, true)
	defer client.Close()
	return c.String(http.StatusOK, id)
}

func postAddPlayer(c echo.Context) error {
	reqBody := new(types.Player)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	roomId := reqBody.RoomId
	playerName := reqBody.PlayerName
	playerIcon := reqBody.PlayerIcon
	playerId := createDB.AddPlayer(roomId, playerName, playerIcon)
	return c.JSON(http.StatusOK, playerId)
}

func postCreateTheme(c echo.Context) error {
	reqBody := new(types.Theme)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	playerId := reqBody.PlayerId
	theme := reqBody.Text
	roomId := reqBody.RoomId
	return c.JSON(http.StatusOK, createDB.CreateTheme(theme, playerId, roomId))
}
func postCreateHint(c echo.Context) error {
	reqBody := new(types.Hint)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	playerId := reqBody.PlayerId
	hint := reqBody.Hint
	roomId := reqBody.RoomId
	return c.JSON(http.StatusOK, createDB.CreateHint(hint, playerId, roomId))
}
func postDeleteHint(c echo.Context) error {
	reqBody := new(types.DeleteHint)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	hintList := reqBody.Hint
	roomId := reqBody.RoomId
	return c.JSON(http.StatusOK, createDB.DeleteHint(hintList, roomId))
}
func postStartGame(c echo.Context) error {
	reqBody := new(types.Game)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	roomId := reqBody.RoomId

	return c.JSON(http.StatusOK, createDB.StartGame(roomId))
}
func postDecideTheme(c echo.Context) error {
	reqBody := new(types.DecideTheme)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	roomId := reqBody.RoomId
	howToDecideTheme := reqBody.HowToDecideTheme
	return c.JSON(http.StatusOK, createDB.DecideTheme(roomId, howToDecideTheme))
}
func postUpdateAnswer(c echo.Context) error {
	reqBody := new(types.Answer)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	playerId := reqBody.PlayerId
	roomId := reqBody.RoomId
	answer := reqBody.Answer

	return c.JSON(http.StatusOK, createDB.UpdateAnswer(answer, roomId, playerId))
}
func postIsCorrect(c echo.Context) error {
	reqBody := new(types.IsCorrect)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	roomId := reqBody.RoomId
	isCorrect := reqBody.IsCorrect

	return c.JSON(http.StatusOK, createDB.IsCorrect(roomId, isCorrect))
}
func postEndGame(c echo.Context) error {
	reqBody := new(types.Game)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	roomId := reqBody.RoomId

	return c.JSON(http.StatusOK, createDB.EndGame(roomId))
}

func postVote(c echo.Context) error {
	reqBody := new(types.Vote)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	playerId := reqBody.PlayerId
	roomId := reqBody.RoomId
	inputPlayerId := reqBody.InputPlayerId
	return c.JSON(http.StatusOK, createDB.Vote(playerId, inputPlayerId, roomId))
}
func postJudgementWolf(c echo.Context) error {
	reqBody := new(types.Vote)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	playerId := reqBody.PlayerId
	roomId := reqBody.RoomId
	return c.JSON(http.StatusOK, readDB.JudgementWolf(roomId, playerId))
}
func postAddStep(c echo.Context) error {
	reqBody := new(types.Game)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	roomId := reqBody.RoomId

	return c.JSON(http.StatusOK, createDB.AddStep(roomId))
}
