package ledger

type LedgerConfig struct {
	Journal journal
}

type journal struct {
	Dir         string
	DefaultPerm int
}
