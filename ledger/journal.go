package ledger

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Journal struct {
	Accounts []string
	Path     string
}

const aliasTemplate = `###START-ALIAS-{{.Name}}
alias {{.Name}} = {{.AccountName}}
###END-ALIAS-{{.Name}}
`

// Get existing account from ledger file
func (j *Journal) GetAccounts() []string {
	out, _ := ExecLedgerCommand(j.Path, "accounts")
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

// Get last transaction ID from ledger file
func (j *Journal) getLastTransactionId() int {
	f, err := os.OpenFile(j.Path, os.O_RDONLY, 0666) // TODO: get perm from config
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

// AddTransaction adds given transaction to ledger file
func (j *Journal) AddTransaction(t *Transaction) error {

	f, err := os.OpenFile(j.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666) // TODO: get perm from config
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// Set transaction ID
	t.ID = j.getLastTransactionId() + 1

	if _, err = f.WriteString(t.Render()); err != nil {
		return err
	}

	return nil
}

func (j *Journal) DeleteAlias(name string) error {
	m, _ := regexp.MatchString(`^\w+$`, name)
	if !m {
		return errors.New(`alias name not matched '^\w+$'`)
	}

	cmd := fmt.Sprintf(`'/###START-ALIAS-%[1]s/,/###END-ALIAS-%[1]s/d'`, name)
	return ExecSedCommandOnFile(j.Path, cmd)
}

func (j *Journal) AddAlias(name, accountName string) error {
	matced, _ := regexp.MatchString(`^\w+$`, name)
	accNameMatched, _ := regexp.MatchString(`^[\w:-]+$`, accountName)
	if !matced || !accNameMatched {
		return errors.New(`invalid alias name or account name format`)
	}

	alias  := struct {
		Name        string
		AccountName string
	}{
		name,
		accountName,
	}

	tmpl, err := template.New("addAlias").Parse(aliasTemplate)

	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, alias)

	if err != nil {
		return err
	}

	return InsertToBeginningOfFile(j.Path, buf.String())
}


func (j *Journal) DeleteTransaction(id string) error {
	m, _ := regexp.MatchString(`^\d+$`, id)
	if !m {
		return errors.New(`transaction id not matched '^\d+$'`)
	}

	cmd := fmt.Sprintf(`'/###START-TRANS:%[1]s/,/###END-TRANS%[1]s/d'`, id)
	return ExecSedCommandOnFile(j.Path, cmd)
}

