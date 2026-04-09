package quiz

// import "database/sql"

type Quiz struct {
	Id        int
	Title     string
	AuthorId  int
	Questions []Question
}

func NewQuiz() *Quiz {
	return &Quiz{}
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
	ImageId     *string
	Answers     []Answer
}

func NewQuestion() *Question {
	return &Question{}
}

type Image struct {
	Id       string
	ImageURL string
}

func NewImage() *Image {
	return &Image{}
}

type Answer struct {
	Id          int
	TextContent string
	IsCorrect   bool
	QuestionId  int
}

func NewAnswer() *Answer {
	return &Answer{}
}
