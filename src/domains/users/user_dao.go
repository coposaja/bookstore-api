package users

import (
	"fmt"

	"github.com/coposaja/bookstore-api/src/db/mysql"
	"github.com/coposaja/bookstore-api/src/utils/rerr"
)

const (
	// UserStatusActive constant indicating user status is active
	UserStatusActive = "active"
)

// Save User to DB
func (user *User) Save() rerr.RestError {
	query, err := mysql.Client.Prepare("INSERT INTO users (FirstName, LastName, Email, DateCreated, Status) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return rerr.NewInternalServerError(err.Error())
	}
	defer query.Close()

	insertResult, err := query.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.DateCreated,
		user.Status,
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
	query, err := mysql.Client.Prepare("SELECT Id, FirstName, LastName, Email, DateCreated, Status FROM users WHERE Id = ?")
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
		&user.Status,
	); err != nil {
		return rerr.NewInternalServerError(fmt.Sprintf("Error while trying to get User Id %d: %s", user.ID, err.Error()))
	}

	defer query.Close()
	return nil
}

// Update User data
func (user *User) Update() rerr.RestError {
	query, err := mysql.Client.Prepare("UPDATE users SET FirstName = ?, LastName = ?, Email = ?, Status = ? WHERE Id = ?")
	if err != nil {
		return rerr.NewInternalServerError(err.Error())
	}
	defer query.Close()

	_, err = query.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.Status,
		user.ID,
	)
	if err != nil {
		return rerr.NewInternalServerError(fmt.Sprintf("Error while trying to update User id %d: %s", user.ID, err.Error()))
	}

	return nil
}

// Delete User data
func (user *User) Delete() rerr.RestError {
	query, err := mysql.Client.Prepare("DELETE FROM users WHERE Id = ?")
	if err != nil {
		return rerr.NewInternalServerError(err.Error())
	}
	defer query.Close()

	_, err = query.Exec(user.ID)
	if err != nil {
		return rerr.NewInternalServerError(fmt.Sprintf("Error while trying to delete User id %d: %s", user.ID, err.Error()))
	}

	return nil
}

// Search User data using certain parameter
func (user *User) Search(status string) ([]User, rerr.RestError) {
	query, err := mysql.Client.Prepare("SELECT Id, FirstName, LastName, Email, DateCreated, Status FROM users WHERE Status = ?")
	if err != nil {
		return nil, rerr.NewInternalServerError(err.Error())
	}
	defer query.Close()

	rows, err := query.Query(status)
	if err != nil {
		return nil, rerr.NewInternalServerError(fmt.Sprintf("Error while trying to find User with status %s: %s", status, err.Error()))
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, rerr.NewInternalServerError(fmt.Sprintf("Error while trying to find User with status %s: %s", status, err.Error()))
		}
		results = append(results, user)
	}

	return results, nil
}
