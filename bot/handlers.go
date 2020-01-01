package bot

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"

	"github.com/mesuutt/teledger/db"
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
	} else {
		result := GetRegexSubMatch(aliasRegex, m.Payload)
		err := db.AddAlias(m.Sender.ID, "foo", "expenses:asdasd")
		if result["name"] == "" || result["accName"] == "" {
			_, err = h.Bot.Send(m.Sender, aliasAddErr, &tb.SendOptions{
				ParseMode: "Markdown",
			})
		}

		fmt.Println(err, result["name"], result["accName"])
	}
}

func (h *Handler) Text(m *tb.Message) {
	if IsAliasCommand(m.Text) {
		cb := confirmAlias
		cb.Data = "confirm:" + m.Text
		inlineBtns := [][]tb.InlineButton{{cb, cancelBtn}}
		msg, err := h.Bot.Send(m.Sender, "Confirm?", &tb.ReplyMarkup{
			InlineKeyboard: inlineBtns,
		})
		fmt.Println(err, msg)
		return
	}

	_, _ = h.Bot.Send(m.Sender, unknownCommandErrMsg)
}

func (h *Handler) Confirm(c *tb.Callback) {
	msg := c.Data[8:] // remove confirm from text
	result := GetRegexSubMatch(aliasRegex, msg)
	err := db.AddAlias(c.Sender.ID, result["name"], result["accName"])
	if err != nil {
		logrus.Error(errors.Wrap(err, "error on confirm: "+c.Data))
		// h.Bot.Send(c.Sender, "An error occurred")
		err := h.Bot.Respond(c, &tb.CallbackResponse{
			Text: "An error occurred",
		})
		fmt.Println(err)
		return
	}

	// FIXME: should be confirmed only once
	_ = h.Bot.Respond(c, &tb.CallbackResponse{
		CallbackID: c.ID,
	})

	// h.Bot.Send(c.Sender, "Alias added")
}

func (h *Handler) Cancel(c *tb.Callback) {
	h.Bot.Send(c.Sender, "Cancelled")
}
