package youcanbookme

import (
	"github.com/dghubble/sling"
)

func GetAuth(email, password string) *sling.Sling {
	s := sling.New().SetBasicAuth(email, password)
	return s
}