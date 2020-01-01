package bot

import (
	"bytes"
	"fmt"
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
var deleteAliasRegex = `del alias (?P<name>\w+)$`

func (h *Handler) Alias(m *tb.Message) {
	if m.Payload == "" {
		_, err := h.Bot.Send(m.Sender, aliasHelp, &tb.SendOptions{
			ParseMode: "Markdown",
		})
		if err != nil {
			logrus.Error(err)
		}
	}
}

func (h *Handler) Text(m *tb.Message) {
	if IsAliasCommand(m.Text) {
		match := GetRegexSubMatch(setAliasRegex, m.Text)
		err := SetAlias(m.Sender.ID, match["name"], match["accName"])
		if err != nil {
			logrus.Error(errors.Wrap(err, "error when set alias: "+m.Text))
			_, err := h.Bot.Send(m.Sender, aliasError)
			if err != nil {
				logrus.Error(err)
			}
			return
		}

		_, err = h.Bot.Send(m.Sender, "Alias added")
		if err != nil {
		    logrus.Error(err)
		}

		return
	}

	if m.Text == "show aliases" {
		aliases := db.GetUserAliases(m.Sender.ID)
		res := new(bytes.Buffer)
		res.WriteString("Alias = AccountName\n")
		res.WriteString("===============\n")

		for k, v := range aliases {
			res.WriteString(fmt.Sprintf("%s = %s\n", k, v))
		}

		h.Bot.Send(m.Sender, res.String())
		return
	}

	if strings.HasPrefix(m.Text, "del alias") {
		match := GetRegexSubMatch(deleteAliasRegex, m.Text)
		if _, ok := match["name"]; !ok {
			h.Bot.Send(m.Sender, "Alias name not matched. \nUsage: "+ delAliasHelp)
			return
		}

		err := DeleteAlias(m.Sender.ID, match["name"])
		if err != nil {
			logrus.Error(errors.Wrap(err, "alias delete error: " + m.Text))
			h.Bot.Send(m.Sender, aliasError)
			return
		}

		h.Bot.Send(m.Sender, "Alias deleted")
		return
	}

	_, _ = h.Bot.Send(m.Sender, unknownCommandErrMsg)
}
