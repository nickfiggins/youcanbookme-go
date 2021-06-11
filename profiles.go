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


func newProfilesService(sling *sling.Sling, user Account) *ProfilesService {
	return &ProfilesService{
		sling: sling.Path("profiles/"),
		CurrentUser: user,
	}
}

func (s *ProfilesService) GetProfiles() ([]Profile, *http.Response, error) {
	apiError := new(APIError)
	var profiles []Profile
	resp, err := s.sling.New().Receive(&profiles, apiError)
	return profiles, resp, relevantError(err, *apiError)
}

func (s *ProfilesService) GetProfileById(id string) (Profile, *http.Response, error) {
	apiError := new(APIError)
	var profile Profile
	resp, err := s.sling.New().Path(id).Receive(&profile, apiError)
	return profile, resp, relevantError(err, *apiError)
}