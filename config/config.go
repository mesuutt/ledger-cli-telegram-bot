package config

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type env struct {
	DBFile    string
	LedgerCLI struct {
		JournalDir string
	}
	Telegram struct {
		Token string
	}
	Logging logging
}

type logging struct {
	Level  string
	Format string
}

var (
	Env env
)

func Init() {
	Env.LedgerCLI.JournalDir = os.Getenv("TELEDGER_JOURNAL_DIR")
	if Env.LedgerCLI.JournalDir == "" {
		logrus.Fatal("TELEDGER_JOURNAL_DIR env var cannot be empty")
	}

	if _, err := os.Stat(Env.LedgerCLI.JournalDir); os.IsNotExist(err) {
		logrus.Fatal(fmt.Sprintf("Ensure journal directory exist and writable: %s", Env.LedgerCLI.JournalDir))
	}

	Env.Telegram.Token = os.Getenv("TELEDGER_TELEGRAM_TOKEN")
	if Env.Telegram.Token == "" {
		logrus.Fatal("TELEDGER_TELEGRAM_TOKEN env cannot be empty")
	}

	Env.DBFile = os.Getenv("TELEDGER_DB_FILE")
	if Env.DBFile == "" {
		logrus.Info("TELEDGER_DB_FILE env not set. `ledger.db` will be used.")
		Env.DBFile = "ledger.db"
	}

}
