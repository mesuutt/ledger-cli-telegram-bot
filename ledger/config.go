package ledger

type Config struct {
	Journal struct {
		Dir         string
		DefaultPerm int
	}
}
