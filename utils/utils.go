package utils

import (
	"math/rand"
	"time"

	"github.com/go-playground/validator"
)

const RandomStringCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func HandleValidatorErrors (err error) map[string]interface{} {
	errs, ok := err.(validator.ValidationErrors)
	respErrors := make(map[string]interface{})

	if ok {
		for _, v := range errs {
			fieldName := v.StructNamespace()
			respErrors[fieldName] = v.Tag()
		}
	} else {
		respErrors["nonFieldError"] = "Invalid data"
	}
	return respErrors
}

func GetRandomString(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = RandomStringCharset[seededRand.Intn(len(RandomStringCharset))]
	}
	return string(b)
}