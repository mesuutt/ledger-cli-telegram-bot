package ledger

type Config struct {
	JournalDir string
}

/*
var conf Config
if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
    // handle error
}
*/
