package schemas

type TodoItem struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	CreatedAt   string
}

type HTTPError struct {
	Code    int
	Message string
}
