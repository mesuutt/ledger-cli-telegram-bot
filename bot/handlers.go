package bot

import (
	"fmt"

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
		cb := confirmAlias
		cb.Data = "confirm:" + m.Text
		inlineBtns := [][]tb.InlineButton{{cb, cancelBtn}}
		msg, err := h.Bot.Send(m.Sender, "Are you sure?", &tb.ReplyMarkup{
			InlineKeyboard: inlineBtns,
		})
		fmt.Println(err, msg)
		return
	}

	_, _ = h.Bot.Send(m.Sender, unknownCommandErrMsg)
}

func (h *Handler) Confirm(c *tb.Callback) {
	msg := c.Data[8:] // remove "confirm" from text

	if IsAliasCommand(msg) {
		match := GetRegexSubMatch(aliasRegex, msg)
		err := SetAlias(c.Sender.ID, match["name"], match["accName"])
		if err != nil {
			logrus.Error(errors.Wrap(err, "error on confirm: "+c.Data))
			err := h.Bot.Respond(c, &tb.CallbackResponse{
				Text: "An error occurred",
			})
			if err != nil {
			    logrus.Error(err)
			}
			return
		}

		h.Bot.Send(c.Sender, "Alias added")
	}

	// FIXME: should be confirmed only once
	_ = h.Bot.Respond(c, &tb.CallbackResponse{
		CallbackID: c.ID,
	})
}

func (h *Handler) Cancel(c *tb.Callback) {
	h.Bot.Send(c.Sender, "Cancelled")
}
