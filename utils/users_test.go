package utils

import (
	"strings"
	"testing"
)

func TestGetRandomString(t *testing.T) {
	randomString := GetRandomString(5)
	if length := len(randomString); length != 5 {
		t.Errorf("getRandomString length was incorrect, got: %d, want: %d", length, 5)
	}

	randomString2 := GetRandomString(5)
	if strings.Compare(randomString, randomString2) == 0 {
		t.Errorf("getRandomString does not return a random value" +
			"first got %s, second got %s", randomString, randomString2)
	}
}