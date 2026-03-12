package store

type Quiz struct {
	Id       int
	Name     string
	AuthorId int
}

func NewQuiz() *Quiz {
	return &Quiz{}
}

type Game struct {
	Id        string
	Code      string
	IsStarted bool
	QuizId    int
}

func NewGame() *Game {
	return &Game{}
}

type User struct {
	Id           int
	Username     string
	PasswordHash string
}

func NewUser() *User {
	return &User{}
}

type Question struct {
	Id          int
	QuizId      int
	Position    int
	TextContent string
	ImageId     int
}

func NewQuestion() *Question {
	return &Question{}
}

type Image struct {
	Id       int
	ImageURL string
}

func NewImage() *Image {
	return &Image{}
}

type Answer struct {
	Id          int
	TextContent string
	IsCorrect   bool
	QuizId      int
}

func NewAnswer() *Answer {
	return &Answer{}
}

type Player struct {
	Id       string  `json:"id"`
	Username string  `json:"username"`
	Score    float32 `json:"score"`
	GameId   string  `json:"game_id"`
}
