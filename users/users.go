package users

import (
	"log"

	"github.com/alexandrevicenzi/unchained"
)


func (user *User) CreateUser(email string, password string) {
	user.Email = normalizeEmail(email)
	user.Password = preparePassword(password)

	user.Save()
}

func (user *User) Save() {
	// TODO: implement this
}

func (user *User) CheckPassword(password string) bool {
	valid, err := unchained.CheckPassword(password, user.Password)

	if err != nil {
		log.Fatal(err)
	}
	return valid
}

