package gotypeform

import "encoding/json"

type User struct {
	Alias    string `json:"alias"`
	Email    string `json:"email"`
	Language string `json:"language"`
	UserID   string `json:"user_id"`
}

func (t *Typeform) GetUser() (*User, error) {
	contents, err := t.buildAndExecRequest("GET", userUrl, nil)
	if err != nil {
		return nil, err
	}
	var user *User
	err = json.Unmarshal(contents, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
