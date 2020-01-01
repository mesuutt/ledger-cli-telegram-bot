package bot

import (
	"gopkg.in/tucnak/telebot.v2"
)

var cancelBtn = telebot.InlineButton{Text: "No", Unique: "cancel_btn", Data: "cancel"}
var confirmAlias = telebot.InlineButton{Text: "Yes", Unique: "confirm_alias", Data: "auto generating"}

