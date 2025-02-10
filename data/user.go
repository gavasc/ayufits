package data

type User struct {
	UserId   int64
	Name     string
	Username string
}

func (u User) Create() error {
	query := "INSERT INTO users (user_id, name, username) VALUES (?, ?, ?);"
	db := ConnectDb()
	if _, err := db.Exec(query, u.UserId, u.Name, u.Username); err != nil {
		return err
	}

	return nil
}
