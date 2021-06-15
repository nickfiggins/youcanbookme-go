package youcanbookme

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrentUserProfiles(t *testing.T) { 

	expectedProfile := Profile{
			ID: "a4e8572c-0a07-4ac4-bebc-bd5ed3d3a8f7",
			AccountId: "720e83b8-1268-4523-9948-3bebfdf0c17c",
			Title: "Test Company",
			Description: "Choose a time that works for you.",
			Subdomain: "testacc1",
			Timezone: "",
			Status: "ONLINE",
	}

	var profiles []Profile

	ycbm := NewTestClient(http.DefaultClient)
	profiles, _, err  := ycbm.CurrentUser.Profiles.GetProfiles(); if err != nil {
		t.Logf("Error fetching current profiles: %v", err)
	}

	expectedProfile.Bookings = profiles[0].Bookings

	assert.Equal(t, 1, len(profiles))
	assert.Equal(t, expectedProfile, profiles[0])
}

func TestUserProfiles(t *testing.T) { 
	ycbm := NewTestClient(http.DefaultClient)
	expectedProfile :=  Profile{
		ID: "a4e8572c-0a07-4ac4-bebc-bd5ed3d3a8f7",
		AccountId: "720e83b8-1268-4523-9948-3bebfdf0c17c",
		Title: "Test Company",
		Description: "Choose a time that works for you.",
		Subdomain: "testacc1",
		Timezone: "",
		Status: "ONLINE",
	}

	userAccount, _, err := ycbm.Accounts.Get(os.Getenv("TEST_ACCOUNT_ID")); if err != nil {
		t.Log("Error fetching user account")
	}

	profiles, _, err  := userAccount.Profiles.GetProfiles(); if err != nil {
		t.Logf("Error fetching current profiles: %v", err)
	}

	expectedProfile.Bookings = profiles[0].Bookings

	assert.Equal(t, 1, len(profiles))
	assert.Equal(t, expectedProfile, profiles[0])
}

func TestUserProfileById(t *testing.T) {
	ycbm := NewTestClient(http.DefaultClient)
	testProfileId := os.Getenv("TEST_PROFILE")
	profile, _, err := ycbm.CurrentUser.Profiles.GetProfileById(testProfileId); if err != nil {
		t.Log("Error finding profile: ", err)
	}

	assert.Equal(t, "Test Company", profile.Title)
}