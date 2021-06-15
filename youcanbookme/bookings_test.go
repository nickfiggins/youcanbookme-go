package youcanbookme
import (
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestCurrentUserBookings(t *testing.T) { 

	ycbm := NewTestClient(http.DefaultClient)

	bookings, _, err := ycbm.CurrentUser.Bookings.GetBookings(); if err != nil {
		t.Fatal(err, bookings)
	}
	if len(bookings) != 0 {
		initBooking := bookings[0]
		assert.Equal(t, ycbm.CurrentUser.ID, initBooking.Accountid)
	}
}

func TestProfileBookings(t *testing.T) { 
	ycbm := NewTestClient(http.DefaultClient)

	userAccount, _, err := ycbm.Accounts.Get(os.Getenv("TEST_ACCOUNT_ID")); if err != nil {
		t.Log("Error fetching user account")
	}

	profiles, _, err  := userAccount.Profiles.GetProfiles(); if err != nil {
		t.Logf("Error fetching current profiles: %v", err)
	}
	
	var bookings [][]Booking
	for _, profile := range profiles {
		curBookings , _ , _ := profile.Bookings.GetBookings()
		bookings = append(bookings, curBookings)
	}
	assert.Equal(t, os.Getenv("TEST_ACCOUNT_ID"), bookings[0][0].Accountid)
}

func TestBookingsById(t *testing.T) {
	ycbm := NewTestClient(http.DefaultClient)
	profile, _, err := ycbm.CurrentUser.Profiles.GetProfileById(os.Getenv("TEST_PROFILE")); if err != nil {
		t.Log("Error finding profile: ", err)
	}

	bookings, _, err :=profile.Bookings.GetBookings(); if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, bookings)
	assert.Equal(t, bookings[0].Accountid, profile.AccountId)

}