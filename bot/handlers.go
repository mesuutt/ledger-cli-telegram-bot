package bot

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Handler struct {
	Bot *tb.Bot
}

var aliasRegex = `alias (?P<name>\w+)\s+(?P<accName>[\w-:]+)$`

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
		match := GetRegexSubMatch(aliasRegex, m.Text)
		err := SetAlias(m.Sender.ID, match["name"], match["accName"])
		if err != nil {
			logrus.Error(errors.Wrap(err, "error when set alias: "+ m.Text))
			_, err := h.Bot.Send(m.Sender, "An error occurred. Please check /help alias")
			if err != nil {
				logrus.Error(err)
			}
			return
		}

		h.Bot.Send(m.Sender, "Alias added")
		return
	}

	_, _ = h.Bot.Send(m.Sender, unknownCommandErrMsg)
}
