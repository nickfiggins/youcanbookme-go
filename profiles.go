package youcanbookme

import (
	"net/http"
	"github.com/dghubble/sling"
)

type Profile struct {
	ID          string      `json:"id"`
	AccountId   string      `json:"accountId"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Subdomain   string      `json:"subdomain"`
	Timezone    string `json:"timeZone"`
	Status      string      `json:"status"`
}

// AccountService provides a method for account credential verification.
type ProfilesService struct {
	sling *sling.Sling
	CurrentUser Account
}

type ProfileInput struct {
	AccountId string
}


func newProfilesService(sling *sling.Sling, user Account) *ProfilesService {
	return &ProfilesService{
		sling: sling.Path("profiles/"),
		CurrentUser: user,
	}
}

func (s *ProfilesService) GetCurrentProfiles() ([]Profile, *http.Response, error) {
	return s.get(ProfileInput{AccountId: s.CurrentUser.ID})
}

func (s *ProfilesService) GetProfiles(accountId string) ([]Profile, *http.Response, error) {
	return s.get(ProfileInput{AccountId: accountId})
}

func (s * ProfilesService) get(ProfileInput) ([]Profile, *http.Response, error) {
	apiError := new(APIError)
	var profiles []Profile
	accountId := s.CurrentUser.ID
	resp, err := s.sling.New().Get("").Path(accountId).Receive(&profiles, apiError)
	return profiles, resp, relevantError(err, *apiError)
}