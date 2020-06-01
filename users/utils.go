package users

import (
	"log"

	"github.com/alexandrevicenzi/unchained"
)

func normalizeEmail(email string) string {
	return email
}

func preparePassword(password string) string {
	hash, err := unchained.MakePassword(password, unchained.GetRandomString(12), "argon2")

	if err != nil {
		log.Fatal(err)
	}

	return hash
}


