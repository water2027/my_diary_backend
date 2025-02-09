package user

import (
	"context"
	"errors"
	"my_diary/database"
	"time"
)

type userHandler interface {
	insert() error
	findUser() bool
	deleteUser() error
	updateUser() error
}

func (user *userModel) insert() error {
	db := database.GetMysqlDB()
	stmt, err := db.Prepare("INSERT INTO users(userEmail, username, password) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.userEmail, user.username, user.password)
	if err != nil {
		return err
	}
	return nil
}

func (user *userModel) findUser() bool {
	db := database.GetMysqlDB()
	var query string
	var value interface{}
	if user.userEmail != "" {
		query = "SELECT * FROM users WHERE userEmail = ?"
		value = user.userEmail
	}
	if user.userId != 0 {
		query = "SELECT * FROM users WHERE userId = ?"
		value = user.userId
	}
	stmt, err := db.Prepare(query)
	if err != nil {
		return false
	}
	defer stmt.Close()
	row := stmt.QueryRow(value)
	err = row.Scan(&user.userId, &user.userEmail, &user.username, &user.password)
	return err == nil
}

func (user *userModel) updateUser() error {
	stmtString := "UPDATE users SET "
	args := []interface{}{}
	hasSet := false
	if user.username != "" {
		stmtString += "username = ?, "
		args = append(args, user.username)
		hasSet = true
	}
	if user.password != "" {
		stmtString += "password = ?, "
		args = append(args, user.password)
		hasSet = true
	}
	if user.userEmail != "" {
		stmtString += "userEmail = ?, "
		args = append(args, user.userEmail)
		hasSet = true
	}
	if !hasSet {
		return errors.New("no update")
	}
	stmtString = stmtString[:len(stmtString)-2]
	stmtString += " WHERE userId = ?"
	args = append(args, user.userId)

	db := database.GetMysqlDB()
	stmt, err := db.Prepare(stmtString)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}
	return nil
}

func (user *userModel) deleteUser() error {
	db := database.GetMysqlDB()
	stmt, err := db.Prepare("DELETE FROM users WHERE userId = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.userId)
	if err != nil {
		return err
	}
	return nil
}

func setCode(ctx context.Context, email, code string, expiry time.Duration) error {
	err := database.SetValue(ctx, email, code, expiry)
	return err
}

func getCode(ctx context.Context, email string) (string, error) {
	code, err := database.GetValue(ctx, email)
	return code, err
}

func deleteCode(ctx context.Context, email string) error {
	return database.DeleteValue(ctx, email)
}