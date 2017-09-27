package ledger

import (
	"bytes"
	"html/template"
	"log"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	FromAccount *Account
	ToAccount   *Account
	Date        string //TODO: Should be time.Time
	Amount      decimal.Decimal
	Payee       string
	Id          int
}

const transactionTemplate = `###START:{{.Id}}
{{.Date}} * {{.Payee}} (##{{.Id}}##)
  {{.ToAccount.Name}}   {{.Amount}}
  {{.FromAccount.Name}}
###END:{{.Id}}
`

func (t *Transaction) String() string {
	tmpl, err := template.New("test").Parse(transactionTemplate)

	if err != nil {
		log.Fatal("Parse: ", err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, t)

	if err != nil {
		panic(err)
	}

	return buf.String()
}
