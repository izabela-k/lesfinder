package users

import (
	"log"
	"time"

	"github.com/alexandrevicenzi/unchained"
	"github.com/google/uuid"
)

func (user *User) New() {
}


func (user *User) CreateUser(email string, password string) {
	user.ID = uuid.New().String()
	now := time.Now()
	user.CreatedAt = &now
	user.Email = normalizeEmail(email)
	user.Password = preparePassword(password)
}


func (user *User) CheckPassword(password string) bool {
	valid, err := unchained.CheckPassword(password, user.Password)

	if err != nil {
		log.Fatal(err)
	}
	return valid
}

