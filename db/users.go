package db

import (
	"anyrom/webservice/types"
	"database/sql"
	"time"
)

func (client *Client) InsertUser(name string, birthDate time.Time) error {
	client.mx.Lock()
	defer client.mx.Unlock()

	_, err := client.Exec(`INSERT INTO users 
		(name, birth_date)
		VALUES ($1, $2)`, name, birthDate)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) GetUser(id int64) (string, time.Time, error) {
	var name sql.NullString
	var birthDate sql.NullTime
	err := client.QueryRow(`SELECT name, birth_date
		FROM users 
		WHERE id = $1`, id).Scan(&name, &birthDate)
	if err != nil {
		return "", time.Now(), err
	}

	return name.String, birthDate.Time, nil
}

func (client *Client) UpdateUser(name string, birthDate time.Time, id int64) error {
	client.mx.Lock()
	defer client.mx.Unlock()

	var count sql.NullInt64
	_ = client.QueryRow(`SELECT count(*)
		FROM users
		WHERE id = $1`, id).Scan(&count)
	if count.Int64 == 0 {
		return types.ErrNoRows
	}

	_, err := client.Exec(`UPDATE users 
		SET name = $1, birth_date = $2 
		WHERE id = $3`, name, birthDate, id)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) DeleteUser(id int64) error {
	_, err := client.Exec(`DELETE
		FROM users
		WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
