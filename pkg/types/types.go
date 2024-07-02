package types

type Book struct {
	Name string `json:"name"`
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
