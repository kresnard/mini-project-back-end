package input

import "time"

type UserInput struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedOn time.Time `json:"created_on"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type FindID struct {
	ID int `json:"user_id"`
}
