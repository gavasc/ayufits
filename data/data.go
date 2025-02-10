package data

import (
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

type FoodType uint

const (
	Fit FoodType = iota + 1
	NaoFit
	Semi
)

func ConnectDb() *sqlx.DB {
	return sqlx.MustConnect("sqlite", "data/data.db")
}

func CreateDB() {
	query := `CREATE TABLE IF NOT EXISTS users (
      user_id integer UNIQUE NOT NULL,
      name text,
      username text);  
      CREATE TABLE IF NOT EXISTS chats (chat_id integer UNIQUE NOT NULL);
      CREATE TABLE IF NOT EXISTS meals (
      user_id integer,
      chat_id integer,
      food_type integer,
      description text,
      timestamp text,
      
      FOREIGN KEY (user_id) REFERENCES users(user_id),
      FOREIGN KEY (chat_id) REFERENCES chats(chat_id));
  `
	db := ConnectDb()
	db.MustExec(query)
}
