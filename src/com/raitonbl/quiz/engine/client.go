package engine

type QuizClient interface {

	IsLast() bool

	IsFirst() bool

	HasNext() bool

	GetScore() int

	GetHint() string

	GetQuestion() string

	AfterPropertiesSet()

	GetMaximumScore() int

	GetOptions() []string

	Answer(answer string) bool

	SetRepository(repository Repository)

}
