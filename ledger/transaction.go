package ledger

import (
	"bytes"
	"html/template"
	"log"
)

type Transaction struct {
	FromAccount *Account
	ToAccount   *Account
	Date        string
	Amount      string
	Payee       string
}

const transactionTemplate = `{{.Date}} * {{.Payee}}
  {{.ToAccount.Name}}   {{.Amount}}
  {{.FromAccount.Name}}
`

func (t *Transaction) String() string {
	tmpl, err := template.New("test").Parse(transactionTemplate)

	if err != nil {
		log.Fatal("Parse: ", err)
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, t)
	return buf.String()
}
