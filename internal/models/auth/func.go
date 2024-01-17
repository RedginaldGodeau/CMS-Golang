package auth

import (
	"cms/internal/database/postgres"
	"cms/internal/privates/structure"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (m UserModel) CheckUser() (bool, error) {
	var sql = `
	SELECT username, password_generator
	FROM api_account
	WHERE username=$1;
`
	conn, err := postgres.Connection()
	if err != nil {
		return false, err
	}

	var buffer UserModel
	err = conn.QueryRow(sql, m.Username).Scan(&buffer.Username, &buffer.Password)
	if err != nil {
		return false, err
	}
	conn.Close()

	err = bcrypt.CompareHashAndPassword([]byte(buffer.Password), []byte(m.Password))

	return err == nil, nil
}

func (m UserModel) Create() error {
	var sql = `
		INSERT INTO 
		api_account(
			username,
			password,
			createAt
		) VALUES (
			 $1,
			 $2,
			 $3
		);
	`

	conn, err := postgres.Connection()
	if err != nil {
		return err
	}

	var querys []any = structure.StructToArray(m)
	querys = append(querys, time.Now().Unix())

	_, err = conn.Exec(sql, querys...)
	if err != nil {
		return err
	}

	return nil
}
