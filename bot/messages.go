package bot


const commands = `
- /help transaction
- /help alias
- /help balance
`

const startMsgFormat = `Hi %s
I am a ledger-cli bot.
You can learn how to use me with:
`  + commands

const unknownCommandErrMsg = `Unknown command :(

` + commands

const setAliasHelp = `
Alias names are case sensitive

*Adding New Alias*:
set alias **AliasName** **AccountName**
a **AliasName** **AccountName**

- Alias names must contains only` + "`" + `[a-zA-Z0-9]` + "`" +
`
- Account names must contains only` + "`" + `[a-zA-Z0-9:_-]` + "`" +
`
*Show existing aliases*:
` + "`a`, `A`, `alias` or `show aliases`"

const aliasHelp = `
You can create aliases and use them as account names instead of writing account names every time when adding transactions.
` + setAliasHelp

const delAliasHelp = `del alias <aliasName>
Alias name must contains only [a-zA-Z0-9_]
`

const delTransactionHelp = `*Delete transaction*:` + "`" + `
del <transactionID>` + "`"

const addTransactionHelp = `
*Adding new transactions*:
There are multiple ways to add new transaction

1: Simple:
` +"`<fromAccount>,<toAccount> <amount> <payee>`"+ `

2: Multiple Account:
` +"`<fromAccount>,<to1>,<to2> <amount> <payee>`"+ `

With this you can add 2 transaction at once, For example:
I transferred money from master card to visa and buy a book from visa.
So I can add 2 transaction shown as below:
`+
"`master,visa,books 123.45`" + `

You can use aliases instead account names.

Also you can write amount using ` + "`qwertyuiop.`" + ` keys (especially useful if you using mobile phone keyboard)
` + "`banka,bankb,alice qwe.rt my debt`" +
`
----

If you want to add historical transaction you can add shown as below:

` +"`<date> <fromAccount>,<toAccount> <amount> <payee>`"+ `

Date format:` +"`dd.MM` or `dd.MM.YYYY`"


const transactionHelp = addTransactionHelp + `

` + delTransactionHelp


const balanceHelp = `
You can show account balance with` + "`balance accountName` command." +
` You can use` + " `b` or `bal` instead `balance`. If you don't give any account name balances of all accounts will be shown."
