package data

import (
	"time"
)

type Meal struct {
	UserId      int64
	ChatId      int64
	FoodType    FoodType `db:"food_type"`
	Description string
	TimeStamp   string
}

func (m Meal) Add() error {
	query := "INSERT INTO meals (user_id, chat_id, food_type, description, timestamp) VALUES (?, ?, ?, ?,?);"
	timestamp := time.Now().Format("2006-01-02 15:04")
	db := ConnectDb()
	defer db.Close()
	if _, err := db.Exec(query, m.UserId, m.ChatId, m.FoodType, m.Description, timestamp); err != nil {
		return err
	}

	return nil
}

func MealInDateRange(userId int64) ([]Meal, error) {
	oneWeekAgo := time.Now().Add(-7 * 24 * time.Hour).Format("2006-01-02 15:04")
	query := "SELECT food_type, description FROM meals WHERE user_id=? AND timestamp BETWEEN ? AND ?;"
	db := ConnectDb()
	defer db.Close()
	meals := []Meal{}

	err := db.Select(&meals, query, userId, oneWeekAgo, time.Now().Format("2006-01-02 15:04"))
	if err != nil {
		return nil, err
	}

	return meals, nil
}
