package data

import "time"

type Meal struct {
	UserId      int64
	ChatId      int64
	FoodType    FoodType
	Description string
	TimeStamp   string
}

func (m Meal) Add() error {
	query := "INSERT INTO meals (user_id, chat_id, food_type, description, timestamp) VALUES (?, ?, ?, ?,?);"
	timestamp := time.Now().Format("2006-01-02 15:04")
	db := ConnectDb()
	if _, err := db.Exec(query, m.UserId, m.ChatId, m.FoodType, m.Description, timestamp); err != nil {
		return err
	}

	return nil
}
