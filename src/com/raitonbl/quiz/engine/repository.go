package engine

type Question struct {
	Text    string
	Hint    string
	Answer  string
	Options []string;
}

type Repository interface {
	Load() []Question;
}