package bot

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"

	"github.com/mesuutt/teledger/db"
)

type Handler struct {
	Bot *tb.Bot
}

var setAliasRegex = `set alias (?P<name>\w+)\s+(?P<accName>[\w-:]+)$`
var deleteAliasRegex = `(del|delete) alias (?P<name>\w+)$`
var deleteTransactionRegex = `(del|delete) (?P<id>\d+)$`

var transactionRegex = `^((?P<day>\d+)\.(?P<month>\d+)(\.(?P<year>\d+)?)\s+)?(?P<from>\w+),(?P<to1>[\w-:]+)(\,(?P<to2>[\w-:]+))?\s+(?P<amount>[\dwqertyuiop.]+)(\s+(?P<payee>.*))?$`

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

func (h *Handler) Help(m *tb.Message) {
	if m.Payload == "alias" {
		_, _ = h.Bot.Send(m.Sender, aliasHelp, &tb.SendOptions{
			ParseMode: "Markdown",
		})
		return
	}

	if m.Payload == "transaction" {
		_, _ = h.Bot.Send(m.Sender, transactionHelp, &tb.SendOptions{
			ParseMode: "Markdown",
		})
		return
	}

	_, err := h.Bot.Send(m.Sender, fmt.Sprintf(startMsgFormat, m.Sender.Username), &tb.SendOptions{
		ParseMode: "Markdown",
	})
	if err != nil {
		logrus.Error(err)
	}
}

func (h *Handler) Text(m *tb.Message) {
	logrus.Info(m.Text)

	if strings.HasPrefix(m.Text, "set alias") {
		match := GetRegexSubMatch(setAliasRegex, m.Text)
		err := SetAlias(m.Sender.ID, match["name"], match["accName"])
		if err != nil {
			logrus.Error(errors.Wrap(err, "error when set alias: "+m.Text))
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

	if m.Text == "a" || m.Text == "alias" || m.Text == "show aliases" {
		aliases := db.GetUserAliases(m.Sender.ID)
		res := new(bytes.Buffer)
		res.WriteString("Alias = AccountName\n")
		res.WriteString("=============\n")

		for k, v := range aliases {
			res.WriteString(fmt.Sprintf("%s = %s\n", k, v))
		}

		h.Bot.Send(m.Sender, res.String())
		return
	}

	if m.Text == "acc" || m.Text == "accounts" || m.Text == "show accounts" {
		res := new(bytes.Buffer)
		res.WriteString("AccountName\n")
		res.WriteString("=============\n")

		for _, v := range GetAccounts(m.Sender.ID) {
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
			logrus.Error(errors.Wrap(err, "alias delete error: "+m.Text))
			_, _ = h.Bot.Send(m.Sender, fmt.Sprintf("Error: %s", err.Error()))
			return
		}

		h.Bot.Send(m.Sender, "Alias deleted")
		return
	}

	if isMatch, err := regexp.Match(deleteTransactionRegex, []byte(m.Text)); isMatch && err == nil {
		match := GetRegexSubMatch(deleteTransactionRegex, m.Text)
		if _, ok := match["id"]; !ok {
			h.Bot.Send(m.Sender, "Invalid delete transaction format.\nUsage: "+delTransactionHelp)
			return
		}

		err := DeleteTransaction(m.Sender.ID, match["id"])
		if err != nil {
			logrus.Error(errors.Wrap(err, "transaction delete error: "+m.Text))
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

}
