package ledger

import (
	"bytes"
	"html/template"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type Transaction struct {
	FromAccount *Account
	ToAccount   *Account
	Date        string //TODO: Should be time.Time
	Amount      decimal.Decimal
	Payee       string
	ID          int
}

const transactionReadTemplate = `
{{.Date}} * {{.Payee}}
  {{.ToAccount.Name}}   {{.Amount}}
  {{.FromAccount.Name}}
`

const transactionWriteTemplate = `###START-TRANS:{{.ID}}
{{.Date}} * {{.Payee}} (##{{.ID}}##)
  {{.ToAccount.Name}}   {{.Amount}}
  {{.FromAccount.Name}}
###END-TRANS:{{.ID}}
`

func (t *Transaction) String() string {
	tmpl, err := template.New("test").Parse(transactionReadTemplate)
	if err != nil {
		logrus.Error("Parse: ", err)
		return "" // FIXME: return error
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, t)

	if err != nil {
		logrus.Error(err)
		return ""
	}

	return buf.String()
}


func (t *Transaction) Render() string {
	tmpl, err := template.New("test").Parse(transactionWriteTemplate)
	if err != nil {
		logrus.Error("Parse: ", err)
		return "" // FIXME: return error
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, t)

	if err != nil {
		logrus.Error(err)
		return ""
	}

	return buf.String()
}
