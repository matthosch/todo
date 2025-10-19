package todo

type todo struct {
	Description string
	Completed   bool
}

func NewTodo(description string) *todo {
	return &todo{
		Description: description,
		Completed:   false,
	}
}
