package commands

import (
	"ayufits/data"
	"fmt"

	"gopkg.in/telebot.v4"
)

func HandleResumo(c telebot.Context) string {
	meals, err := data.MealInDateRange(c.Sender().ID)
	if err != nil {
		return err.Error()
	}

	ftypeCount := countFoodTypes(meals)
	res := resumoResponse(ftypeCount)

	return res
}

func countFoodTypes(meals []data.Meal) map[data.FoodType]uint {
	counts := map[data.FoodType]uint{data.Fit: 0, data.NaoFit: 0, data.Semi: 0}

	for _, meal := range meals {
		counts[meal.FoodType] += 1
	}

	return counts
}

func resumoResponse(counts map[data.FoodType]uint) string {
	return fmt.Sprintf("Nessa ultima semana você teve um total de:\n%d refeições fit\n%d refeições não fit\n%d refeições semi fit",
		counts[data.Fit], counts[data.NaoFit], counts[data.Semi])
}
