package users

import (
	"time"
)

type User struct {
	Id string
	Email string
	Password string

	CreatedAt time.Time
	IsConfirmed bool
	IsAdmin bool
	IsBlocked bool
	IsCompleted bool

	Name string
	BirthDate time.Time
	Description string

	Photos []Photo

	Likes []Like
	Liked []Like

	Location
	City string
	LocationHilbertA string
	LocationHilbertB string
}

type Location struct {
	lat float64
	lng float64
}

type Like struct {
	CreatedAt time.Time
	User
}

type Photo struct {
	Id string
	User
	IsAvailable bool
	OriginalFile string

	Variants []PhotoVariant
}

type PhotoVariant struct {
	File string
	height int
	width int
}
