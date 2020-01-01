package bot


const commands = `
- /help add
- /help alias
- /help account
- /help export
`

const startMsgFormat = `Hi %s
I am a ledger-cli bot.
You can learn how to use me with:
`  + commands

const unknownCommandErrMsg = `Unknown command :(

` + commands

const setAliasHelp = `
*Show existing aliases*:
show aliases

*Adding New Alias*:
set alias **AliasName** **AccountName**

- Aliases must contains only` + "`" + `[a-zA-Z0-9]` + "`" +
`
- Account names must contains only` + "`" + `[a-zA-Z0-9:_-]` + "`" +
`
`

const aliasHelp = `
You can create aliases and use them as account names instead of writing account names every time when adding transactions.
` + setAliasHelp

const delAliasHelp = `del alias aliasName
Alias name must contains only [a-zA-Z0-9_]
`