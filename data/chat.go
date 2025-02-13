package data

type Chat struct {
	ChatId int64
}

func (c Chat) Create() error {
	query := "INSERT INTO chats (chat_id) VALUES (?);"
	db := ConnectDb()
	defer db.Close()
	if _, err := db.Exec(query, c.ChatId); err != nil {
		return err
	}

	return nil
}
