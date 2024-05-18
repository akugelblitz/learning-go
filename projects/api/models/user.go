package models

import (
	"errors"
	"fmt"

	"example.com/api/db"
	"example.com/api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := `
  INSERT INTO users(email, password) VALUES (?, ?)
  `

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()

	hasedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		fmt.Println(err)
		return err
	}

	result, err := stmt.Exec(user.Email, hasedPassword)
	if err != nil {
		fmt.Println(err)
		return err
	}
	id, err := result.LastInsertId()

	user.ID = id
	fmt.Println(err)
	return err
}

func (user User) ValidateCreds() error {
	query := `
  SELECT id, password from users WHERE email = ?
  `
	row := db.DB.QueryRow(query, user.Email)
	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)
	if err != nil {
		return errors.New("Credentials invalid")
	}
	fmt.Println("password: ", user.Password)
	passwordOk := utils.IsPasswordCorrect(user.Password, retrievedPassword)

	if !passwordOk {
		return errors.New("Credentials invalid")
	}

	return nil
}
