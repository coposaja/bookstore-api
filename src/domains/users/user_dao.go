package users

import (
	"fmt"

	"github.com/coposaja/bookstore-api/src/db/mysql"
	"github.com/coposaja/bookstore-api/src/utils/date"
	"github.com/coposaja/bookstore-api/src/utils/rerr"
)

// Save User to DB
func (user *User) Save() rerr.RestError {
	query, err := mysql.Client.Prepare("INSERT INTO users (FirstName, LastName, Email, DateCreated) VALUES (?, ?, ?, ?)")
	if err != nil {
		return rerr.NewInternalServerError(err.Error())
	}
	defer query.Close()
	user.DateCreated = date.GetNow()

	insertResult, err := query.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.DateCreated,
	)
	if err != nil {
		return rerr.NewInternalServerError(fmt.Sprintf("Error while trying to save User: %s", err.Error()))
	}

	uid, err := insertResult.LastInsertId()
	if err != nil {
		return rerr.NewInternalServerError(fmt.Sprintf("Error while trying to save User: %s", err.Error()))
	}

	user.ID = int(uid)
	return nil
}

// Get User from DB by UserID
func (user *User) Get() rerr.RestError {
	query, err := mysql.Client.Prepare("SELECT Id, FirstName, LastName, Email, DateCreated FROM users WHERE Id = ?")
	if err != nil {
		return rerr.NewInternalServerError(err.Error())
	}

	result := query.QueryRow(user.ID)
	if err := result.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.DateCreated,
	); err != nil {
		return rerr.NewInternalServerError(fmt.Sprintf("Error while trying to get User Id %d: %s", user.ID, err.Error()))
	}

	defer query.Close()
	return nil
}

// Update User data
func (user *User) Update() rerr.RestError {
	query, err := mysql.Client.Prepare("UPDATE users SET FirstName = ?, LastName = ?, Email = ? WHERE Id = ?")
	if err != nil {
		return rerr.NewInternalServerError(err.Error())
	}
	defer query.Close()

	_, err = query.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.ID,
	)
	if err != nil {
		return rerr.NewInternalServerError(fmt.Sprintf("Error while trying to update User id %d: %s", user.ID, err.Error()))
	}

	return nil
}
