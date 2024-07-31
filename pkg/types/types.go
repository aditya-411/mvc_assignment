package types

type Book struct {
	Name string `json:"name"`
	Author string `json:"author"`
	Publisher string `json:"publisher"`
	Quantity string `json:"quantity"`
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
	Id int
	Username string
	Name string
	Author string
	Type string
	Fine int
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

type BookCatalogue struct {
	Message string
	Books []Book
}

type BookRequestManagementPage struct {
	Message string
	Requests []PendingApproval
}

type AdminAccessRequestsPage struct {
	Message string
	Requests []User
}