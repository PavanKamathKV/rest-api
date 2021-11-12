// model.go

package main

import (
	"database/sql"
	"fmt"
)

type user struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	Email     string `josn:"email"`
}

func (u *user) getUser(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT firstname, lastname, age, email FROM users WHERE id=%d", u.ID)
	return db.QueryRow(statement).Scan(&u.FirstName, &u.LastName, &u.Age, &u.Email)
}

func (u *user) updateUser(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE users SET firstname='%s', lastname='%s', age=%d,email='%s' WHERE id=%d", u.FirstName, u.LastName, u.Age, u.Email, u.ID)
	_, err := db.Exec(statement)
	return err
}

func (u *user) deleteUser(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM users WHERE id=%d", u.ID)
	_, err := db.Exec(statement)
	return err
}

func (u *user) createUser(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO users(firstname, lastname, age, email) VALUES('%s', '%s', %d, '%s')", u.FirstName, u.LastName, u.Age, u.Email)
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}

func getUsers(db *sql.DB, start, count int) ([]user, error) {
	statement := fmt.Sprintf("SELECT id, firstname, lastname, age, email FROM users LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []user{}

	for rows.Next() {
		var u user
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
