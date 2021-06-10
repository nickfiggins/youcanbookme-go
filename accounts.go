package youcanbookme

import (
	"net/http"
	"time"
	"github.com/dghubble/sling"
)
	
type Account struct {
	ID                string    `json:"id"`
	Email             string    `json:"email"`
	Trialendsat       time.Time `json:"trialEndsAt"`
	Quantitypaidfor   int       `json:"quantityPaidFor"`
	Quantityforfree   int       `json:"quantityForFree"`
	Quantityfreetrial int       `json:"quantityFreeTrial"`
	Quantityallocated int       `json:"quantityAllocated"`
	Currency          string    `json:"currency"`
	Plan              string    `json:"plan"`
	Planmonths        int       `json:"planMonths"`
	Smscredits        int       `json:"smsCredits"`
}

// AccountService provides a method for account credential verification.
type AccountsService struct {
	sling *sling.Sling
}


func newAccountsService(sling *sling.Sling) *AccountsService {
	return &AccountsService{
		sling: sling,
	}
}

func (s *AccountsService) Get(id string) (Account, *http.Response, error) {
	apiError := new(APIError)
	var acc Account
	resp, err := s.sling.New().Get("").Path(id).Receive(&acc, apiError)
	return acc, resp, relevantError(err, *apiError)
}

func (s *AccountsService) GetCurrentAccount() (Account, *http.Response, error) {
	return GetSelf(s.sling)
}