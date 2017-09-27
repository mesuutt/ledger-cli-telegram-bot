package ledger

import (
	"fmt"
	"os"
	"path"
)

type User struct {
	Username string
	Journal  Journal
}

func (user *User) GetAccounts() []string {
	journal := user.GetJournal()
	return journal.GetAccounts()
}

func (user *User) GetJournal() Journal {
	dataDir := os.Getenv("LEDGER_DATA_DIR")
	if dataDir == "" {
		dataDir = "./data"
	}

	return Journal{Path: path.Join(dataDir, fmt.Sprintf("%s.dat", user.Username))}
}
