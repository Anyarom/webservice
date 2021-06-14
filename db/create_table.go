package db

func (client *Client) CreateTable() error {
	_, err := client.Exec(
		"CREATE TABLE IF NOT EXISTS users (id INTEGER constraint users_pk primary key autoincrement, name STRING, birth_date DATE);" +
			"CREATE unique INDEX IF NOT EXISTS users_id_uindex on users (id)")
	if err != nil {
		return err
	}
	return nil
}
