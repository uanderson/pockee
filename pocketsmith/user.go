package pocketsmith

// User contains the information regarding
// the logged user on PocketSmith.
type User struct {
	Id int64
}

func fetchUser() (*User, error) {
	var user *User

	err := FetchJson("/me", &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
