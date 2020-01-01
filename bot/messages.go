package bot

const unknownCommandErrMsg = `Unknown command :(

/help
`
const setAliasHelp = `*Adding New Alias*:
set alias **AliasName** **AccountName**

- Aliases must contains only` + "`" + `[a-zA-Z0-9]` + "`" +
`
- Account names must contains only` + "`" + `[a-zA-Z0-9:_-]` + "`"

const aliasHelp = `
You can use aliases instead of write account name everytime

` + setAliasHelp
