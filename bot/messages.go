package bot


const commands = `
- /help transaction
- /help alias
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

- Aliases must contains only` + "`" + `[a-zA-Z0-9]` + "`" +
`
- Account names must contains only` + "`" + `[a-zA-Z0-9:_-]` + "`" +
`
*Show existing aliases*:
` + "`a` or `alias` or `show aliases`"

const aliasHelp = `
You can create aliases and use them as account names instead of writing account names every time when adding transactions.
` + setAliasHelp

const delAliasHelp = `del alias <aliasName>
Alias name must contains only [a-zA-Z0-9_]
`

const delTransactionHelp = `
*Delete transaction*:` + "`" + `
del <transactionID>` + "`"

const addTransactionHelp = `
*Adding new transactions*:
There are multiple ways to add new transaction

1: Simple:
` +"`" + `<fromAccount>,<toAccount> <amount> <desc>` + "`"+ `

2: Multiple Account:
` +"`" + `<fromAccount>,<to1>,<to2> <amount> <desc>` + "`"+ `

With this you can add 2 transaction at once, For example:
I transferred money from A bank to B bank and send the money to my friend from B bank.
So I can add 2 transaction shown as below:
`+
"`" + `banka,bankb,alice 123.45` + "`" + `
Also you can write amount using ` + "`"+ `qwertyuiop.` + "`" + ` keys (especially useful if you using mobile phone keyboard)
` + "`" + `banka,bankb,alice qwe.rt my debt` + "`" +
`
----

If you want to add historical transaction you can add shown as below:

` +"`" + `<date> <fromAccount>,<toAccount> <amount> <desc>` + "`"+ `

Date format:` +"`" + `dd.MM` + "` or "+" `" + `dd.MM.YYYY` + "`"


const transactionHelp = addTransactionHelp + `

` + delTransactionHelp