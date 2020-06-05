package users

import "time"

type User struct {
	Id string `db:"id"`
	Email string
	Password string

	CreatedAt *time.Time
	IsConfirmed bool
	IsAdmin bool
	IsBlocked bool
	IsCompleted bool

	Name string `dynamodbav:"omitempty"`
	BirthDate *time.Time `dynamodbav:"omitempty"`
	Description string `dynamodbav:"omitempty"`
}

type Token struct {
	Token string
	UserId string
}