package youcanbookme
import (
	"fmt"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCurrentUserProfiles(t *testing.T) { 

	expectedProfile := Profile{
		ID: "198a48fa-1f74-4b61-a093-8be5e8f19474",
		AccountId: "17f60ab5-958a-4c27-89e7-750801c03df8",
		Title: "Nick Figgins",
		Description: "More info: https://nickfiggins.com/tutoring \n\nHi there! Thanks for your interest in scheduling a tutoring session. Feel free to select a time slot below to schedule a meeting.",
		Subdomain: "nickfiggins",
		Timezone: "",
		Status: "ONLINE",
	}

	var profiles []Profile

	ycbm := NewClient(http.DefaultClient)
	fmt.Println(ycbm.CurrentUser.Sling)
	profiles, _, err  := ycbm.CurrentUser.Profiles.GetProfiles(); if err != nil {
		t.Logf("Error fetching current profiles: %v", err)
	}

	assert.Equal(t, 1, len(profiles))
	assert.Equal(t, expectedProfile, profiles[0])
}

func TestUserProfiles(t *testing.T) { 
	ycbm := NewClient(http.DefaultClient)

	expectedProfile := Profile{
			ID: "198a48fa-1f74-4b61-a093-8be5e8f19474",
			AccountId: "17f60ab5-958a-4c27-89e7-750801c03df8",
			Title: "Nick Figgins",
			Description: "More info: https://nickfiggins.com/tutoring \n\nHi there! Thanks for your interest in scheduling a tutoring session. Feel free to select a time slot below to schedule a meeting.",
			Subdomain: "nickfiggins",
			Timezone: "",
			Status: "ONLINE",
	}

	userAccount, _, err := ycbm.Accounts.Get("17f60ab5-958a-4c27-89e7-750801c03df8"); if err != nil {
		t.Log("Error fetching user account")
	}

	profiles, _, err  := userAccount.Profiles.GetProfiles(); if err != nil {
		t.Logf("Error fetching current profiles: %v", err)
	}

	assert.Equal(t, 1, len(profiles))
	assert.Equal(t, expectedProfile, profiles[0])
}

func TestUserProfileById(t *testing.T) {
	ycbm := NewClient(http.DefaultClient)
	profile, _, err := ycbm.CurrentUser.Profiles.GetProfileById("198a48fa-1f74-4b61-a093-8be5e8f19474"); if err != nil {
		t.Log("Error finding profile: ", err)
	}

	assert.Equal(t, "Nick Figgins", profile.Title)
}