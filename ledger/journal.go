package ledger

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Journal struct {
	Accounts []string
	Path     string
}

// Get existing account from ledger file
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

// Get last transaction Id from ledger file
func (j *Journal) getLastTransactionId() int {
	f, err := os.OpenFile(j.Path, os.O_RDONLY, 0666) //TODO: get perm from config
	if err != nil {
		panic(err)
	}

	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		fmt.Println(err)
	}

	buf := make([]byte, 10)
	n, err := f.ReadAt(buf, fi.Size()-int64(len(buf)))
	if err != nil {
		fmt.Println(err)
	}
	buf = buf[:n]
	re := regexp.MustCompile(`###END:(\d+)`)
	match := re.FindStringSubmatch(string(buf))

	if len(match) <= 1 {
		return 0
	}

	i, _ := strconv.Atoi(match[1])

	return i
}

// Add transaction to ledger file
func (j *Journal) AddTransaction(t *Transaction) {

	f, err := os.OpenFile(j.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666) //TODO: get perm from config
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// Set transaction Id
	t.Id = j.getLastTransactionId() + 1

	if _, err = f.WriteString(t.String()); err != nil {
		panic(err)
	}
}
