package types

type Room struct {
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
	RoomId   string `json:"roomId"`
	PlayerId string `json:"playerId"`
	Answer   string `json:"answer"`
}
type IsCorrect struct {
	RoomId    string `json:"roomId"`
	IsCorrect bool   `json:"isCorrect"`
}
type Vote struct {
	InputPlayerId string `json:"inputPlayerId"`
	PlayerId      string `json:"playerId"`
	RoomId        string `json:"roomId"`
}

type RoomPlayer struct {
	RoomId   string
	PlayerId string
}

type CreatePlayer struct {
	PlayerName string
	Icon       int
	Role       int
	Theme      string
	Hint       string
	IsDelete   bool
	Answer     string
	Wolf       bool
	Vote       int
	Point      int
}

type CreateRoom struct {
	PaticNum      int
	Theme         string
	Phase         int
	Step          int
	IsModeWolf    bool
	IsExitWolf    bool
	PeaceVote     int
	IsCorrectWolf bool
	Result        int
}
