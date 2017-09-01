package ledger

import "fmt"

type User struct {
	Username string
	Journal  Journal
}

func (user *User) GetAccounts() []string {
	journal := user.GetJournal()
	return journal.GetAccounts()
}

func (user *User) GetJournal() Journal {
	return Journal{Path: fmt.Sprintf("./data/%s.dat", user.Username)}
}
