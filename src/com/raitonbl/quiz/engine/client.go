package engine

type QuizClient interface {

	GetHint() string

	GetQuestion() string

	GetOptions() []string

	Answer(answer string) bool

	GetScore() int

	HasNext() bool

	IsLast() bool

	IsFirst() bool

	AfterPropertiesSet()

	SetRepository(repository Repository)

}