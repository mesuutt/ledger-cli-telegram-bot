package ledger

import (
	"fmt"
	"path"

	"github.com/mesuutt/teledger/config"
)

type User struct {
	Username string
	Journal  Journal
}

func (u *User) GetAccounts() []string {
	journal := u.GetJournal()
	return journal.GetAccounts()
}
func (u *User) GetJournal() Journal {
	return Journal{Path: path.Join(config.Env.LedgerCLI.JournalDir, fmt.Sprintf("%s.dat", u.Username))}
}

func (u *User) DeleteAlias(name string) error {
	j := u.GetJournal()
	j.DeleteAlias(name)

	return nil
}
func (u *User) AddAlias(name, accountName string) error {
	j := u.GetJournal()
	j.AddAlias(name, accountName)
	return nil
}

func (u *User) DeleteTransaction(id string) error {
	j := u.GetJournal()
	return j.DeleteTransaction(id)
}

func (u *User) GetAccountBalance(accountName string) string {
	j := u.GetJournal()
	return j.GetAccountBalance(accountName)
}

