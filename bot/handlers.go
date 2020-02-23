package bot

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"

	"github.com/mesuutt/teledger/db"
)

type Handler struct {
	Bot *tb.Bot
}

var helpRegex = `^(?i)(h(elp)?)(\s+(?P<name>\w+))?$`
var setAliasRegex = `^(?i)(a(lias)?)\s+(?P<name>\w+)+[\s=]+(?P<accName>[\w-:]+)$`
var showAliasRegex = `^(?i)(show )?a(lias(es)?)?$`
var showAccountsRegex = `^(?i)(show )?acc(ounts?)?$`
var deleteAliasRegex = `^(?i)(del(ete)?) alias (?P<name>\w+)$`
var deleteTransactionRegex = `^(?i)(d(el(ete)?)?) (?P<id>\d+)$`
var showAccountBalanceRegex = `^(?i)(b(al(ance)?)?)(\s+(?P<name>[\w-:]+))?$`

var transactionRegex = `^(((?P<day>\d+)\.(?P<month>\d+)(\.(?P<year>\d+))?)\s+)?(?P<from>[\w:]+),(\s+)?(?P<to1>[\w-:]+)(\,(\s+)?(?P<to2>[\w-:]+))?\s+(?P<amount>[\dwqertyuiop.]+)(\s+(?P<payee>.*))?$`

func (h *Handler) Text(m *tb.Message) {
	logrus.WithField("sender", m.Sender.ID).Info(m.Text)

	if IsRegexMatch(helpRegex, m.Text) {
		match := GetRegexSubMatch(helpRegex, m.Text)
		switch v := match["name"]; v {
		case "alias":
			_, _ = h.Bot.Send(m.Sender, aliasHelp, &tb.SendOptions{
				ParseMode: "Markdown",
			})
		case "transaction":
			_, _ = h.Bot.Send(m.Sender, transactionHelp, &tb.SendOptions{
				ParseMode: "Markdown",
			})
		case "balance":
			_, _ = h.Bot.Send(m.Sender, balanceHelp, &tb.SendOptions{
				ParseMode: "Markdown",
			})
		default:
			_, _ = h.Bot.Send(m.Sender, fmt.Sprintf(startMsgFormat, m.Sender.Username), &tb.SendOptions{
				ParseMode: "Markdown",
			})
		}

		return
	}

	if IsRegexMatch(setAliasRegex, m.Text) {
		match := GetRegexSubMatch(setAliasRegex, m.Text)
		err := SetAlias(m.Sender.ID, match["name"], match["accName"])
		if err != nil {
			logrus.Error(fmt.Errorf("error when set alias: %s, %w", m.Text, err))
			_, err := h.Bot.Send(m.Sender, fmt.Sprintf("Error: %s", err))
			if err != nil {
				logrus.Error(err)
			}
			return
		}

		_, err = h.Bot.Send(m.Sender, "Alias added.")
		if err != nil {
			logrus.Error(err)
		}

		return
	}

	if IsRegexMatch(showAccountBalanceRegex, m.Text) {
		match := GetRegexSubMatch(showAccountBalanceRegex, m.Text)

		// Use original account name if alias exist with given name
		accountName := db.GetAccountByAlias(m.Sender.ID, match["name"])
		if accountName == "" {
			accountName = match["name"]
		}
		bal := GetAccountBalance(m.Sender.ID, accountName)
		h.Bot.Send(m.Sender, bal)
		return
	}

	if IsRegexMatch(showAliasRegex, m.Text) {
		aliases, err := db.GetUserAliases(m.Sender.ID)
		if err != nil {
			if errors.Is(err, &db.ErrBudgetNotFound{}) {
				h.Bot.Send(m.Sender, "budget not found. Please send /start command for create missing budget")
			}
			return
		}

		if len(aliases) == 0 {
			h.Bot.Send(m.Sender, "No alias found")
			return
		}

		res := new(bytes.Buffer)
		res.WriteString("Alias = AccountName\n")
		res.WriteString("=============\n")

		for i, _ := range aliases {
			res.WriteString(fmt.Sprintf("%s = %s\n", aliases[i][0], aliases[i][1]))
		}

		h.Bot.Send(m.Sender, res.String())
		return
	}

	if IsRegexMatch(showAccountsRegex, m.Text) {
		accounts := GetAccounts(m.Sender.ID)
		if len(accounts) == 0 {
			h.Bot.Send(m.Sender, "No account found")
			return
		}

		res := new(bytes.Buffer)
		res.WriteString("AccountName\n")
		res.WriteString("=============\n")

		for _, v := range accounts {
			res.WriteString(fmt.Sprintf("%s\n", v))
		}

		h.Bot.Send(m.Sender, res.String())
		return
	}

	if strings.HasPrefix(m.Text, "del alias") || strings.HasPrefix(m.Text, "delete alias") {
		match := GetRegexSubMatch(deleteAliasRegex, m.Text)
		if _, ok := match["name"]; !ok {
			h.Bot.Send(m.Sender, "Invalid alias name format.\nUsage: "+delAliasHelp)
			return
		}

		err := DeleteAlias(m.Sender.ID, match["name"])
		if err != nil {
			logrus.Error(fmt.Errorf("alias delete error: %s, %w", m.Text, err))
			_, _ = h.Bot.Send(m.Sender, fmt.Sprintf("Error: %s", err.Error()))
			return
		}

		h.Bot.Send(m.Sender, "Alias deleted")
		return
	}

	if isMatch, err := regexp.Match(deleteTransactionRegex, []byte(m.Text)); isMatch && err == nil {
		match := GetRegexSubMatch(deleteTransactionRegex, m.Text)
		err := DeleteTransaction(m.Sender.ID, match["id"])
		if err != nil {
			logrus.Error(fmt.Errorf("transaction delete error: %s, %w", m.Text, err))
			_, _ = h.Bot.Send(m.Sender, fmt.Sprintf("Error: %s", err.Error()))
			return
		}

		h.Bot.Send(m.Sender, "Transaction deleted")
		return
	}

	transactions, err := AddTransaction(m.Sender.ID, m.Text)
	if err != nil {
		_, _ = h.Bot.Send(m.Sender, err.Error())
		_, _ = h.Bot.Send(m.Sender, commands)

		return
	}

	for i, _ := range transactions {
		_, _ = h.Bot.Send(m.Sender, fmt.Sprintf("Transaction added (ID: %d)", transactions[i].ID))
		_, _ = h.Bot.Send(m.Sender, transactions[i].String())
	}

	buf := new(bytes.Buffer)
	buf.Write([]byte("New Balances\n"))
	buf.Write([]byte("===========\n"))
	for i, _ := range transactions {
		bal := GetAccountBalance(m.Sender.ID, transactions[i].FromAccount.Name)
		buf.Write([]byte(bal + "\n"))
		bal = GetAccountBalance(m.Sender.ID, transactions[i].ToAccount.Name)
		buf.Write([]byte(bal + "\n"))
	}
	_, _ = h.Bot.Send(m.Sender, buf.String())

}

func (h *Handler) Start(m *tb.Message) {
	if m.Payload == "" {
		err := db.CreateUser(m.Sender.ID)
		if err != nil {
			logrus.Error(err)
			_, _ = h.Bot.Send(m.Sender, fmt.Sprintf("Account creation failed: %s", err))
			return
		}

		_, err = h.Bot.Send(m.Sender, fmt.Sprintf(startMsgFormat, m.Sender.Username), &tb.SendOptions{
			ParseMode: "Markdown",
		})

		if err != nil {
			logrus.Error(err)
		}
	}
}
