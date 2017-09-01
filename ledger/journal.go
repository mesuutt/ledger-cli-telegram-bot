package ledger

import (
	"os"
	"strings"
)

type Journal struct {
	Accounts []string
	Path     string
}

func (file *Journal) GetPath() string {
	// TODO: @mesut get data dir path with CLI flags.
	//return fmt.Sprintf("./data/%s.dat", file.Path)
	return "./data/%s.dat"
}

func (j *Journal) GetAccounts() []string {
	utils := Utils{}

	out, _ := utils.ExecLedgerCommand(j.Path, "accounts")
	accounts := []string{}

	for {
		line, err := out.ReadString('\n')
		if err != nil {
			break
		}

		// fmt.Printf("LINE: %v", line)
		accounts = append(accounts, strings.TrimRight(line, "\n"))
	}

	// fmt.Printf("%v", out)
	return accounts
}

func (j *Journal) AddTransaction(t *Transaction) {

	f, err := os.OpenFile(j.Path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(t.String()); err != nil {
		panic(err)
	}
}
