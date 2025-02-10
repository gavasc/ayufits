package commands

import (
	"ayufits/data"
	"fmt"

	"gopkg.in/telebot.v4"
)

func HandleStart(c telebot.Context) string {
	var msg string
	chat := data.Chat{
		ChatId: c.Chat().ID,
	}
	if err := chat.Create(); err != nil {
		fmt.Println(err)
		msg += "Chat já cadastrado\n"
	}

	user := data.User{
		UserId:   c.Sender().ID,
		Name:     c.Sender().FirstName,
		Username: c.Sender().Username,
	}
	if err := user.Create(); err != nil {
		fmt.Println(err)
		msg += "Usuário já cadastrado\n"
	} else {
		msg += "Usuário adicionado com sucesso"
	}

	return msg
}
