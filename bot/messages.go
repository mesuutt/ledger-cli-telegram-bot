package bot

const unknownCommandErrMsg = `Unknown command :(

/help
`
const aliasAddHelp = `*Adding New Alias*:
/alias add **aliasName** **AccountName**

- Aliases must contains only` + "`" + `[a-zA-Z0-9]` + "`" +
`
- Account names must contains only` + "`" + `[a-zA-Z0-9:_-]` + "`"


const aliasHelp = `
You can manage ledger-cli account aliases with /alias command

` + aliasAddHelp

const aliasAddErr = `
*Invalid add systax*
---
` + aliasAddHelp