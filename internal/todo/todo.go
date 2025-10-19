package todo

type Todo struct {
	Description string
	Completed   bool
}

func NewTodo(description string) *Todo {
	return &Todo{
		Description: description,
		Completed:   false,
	}
}
