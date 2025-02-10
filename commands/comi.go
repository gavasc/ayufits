package commands

import (
	"ayufits/data"
	"fmt"
	"strings"

	"gopkg.in/telebot.v4"
)

func HandleComi(c telebot.Context) string {
	if valid, str := validComi(c); !valid {
		return str
	}

	msgText := strings.Split(c.Text()[6:], "-")

	food := strings.TrimSpace(msgText[0])
	ftype_str := strings.TrimSpace(msgText[1])
	var ftype data.FoodType

	switch ftype_str {
	case "fit":
		ftype = data.Fit
	case "nao fit", "não fit":
		ftype = data.NaoFit
	case "semi", "semi fit", "semifit":
		ftype = data.Semi
	default:
		return "Tipo de refeição inválido!"
	}

	meal := data.Meal{
		UserId:      c.Sender().ID,
		ChatId:      c.Chat().ID,
		FoodType:    ftype,
		Description: food,
	}
	if err := meal.Add(); err != nil {
		return err.Error()
	}

	return fmt.Sprintf("Comida %s registrada!", ftype_str)
}

func validComi(c telebot.Context) (bool, string) {
	if len(strings.TrimSpace(c.Text())) == 5 {
		return false, "Comando vazio!"
	}

	msgText := strings.Split(c.Text()[6:], "-")
	if len(msgText) < 2 || msgText[1] == "" {
		return false, "Comando faltando argumentos!"
	}

	return true, ""
}
