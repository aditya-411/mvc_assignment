package types

type Book struct {
	Name string `json:"name"`
	Author string `json:"author"`
	Publisher string `json:"publisher"`
}

type ListBooks struct {
	Books []Book `json:"books"`
}

type Message struct {
	Text string
}

type User struct {
	Username string
	IsAdmin  bool
}

type Key string

type IssueBookPage struct {
	Book Book
	Message string
	Show_button bool
}
