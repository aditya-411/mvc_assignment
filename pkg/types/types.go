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
	Message string
}

type Key string

type IssueBookPage struct {
	Book Book
	Message string
	Show_button bool
}

type PendingApproval struct {
	Name string
	Author string
	Type string
}

type PrevTransaction struct {
	Id int
	Name string
	Author string
	IssueDate string
	ReturnDate string
}

type CurrentTransaction struct {
	Id int
	Name string
	Author string
	IssueDate string
	ReturnDate string
	Fine string
}

type UserBooksPage struct {
	Message string
	CurrentlyIssuedBooks []CurrentTransaction
	PendingApprovals []PendingApproval
	PrevTransactions []PrevTransaction
}
