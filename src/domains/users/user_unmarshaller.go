package users

import "encoding/json"

// Marshall User struct to secure User data
func (user User) Marshall(isPrivate bool) interface{} {
	jsonUser, _ := json.Marshal(user)
	if isPrivate {
		var privateUser PrivateUser
		json.Unmarshal(jsonUser, &privateUser)
		return privateUser
	}

	var publicUser PublicUser
	json.Unmarshal(jsonUser, &publicUser)
	return publicUser
}

// Marshall User struct to secure User data
func (users Users) Marshall(isPrivate bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPrivate)
	}

	return result
}
