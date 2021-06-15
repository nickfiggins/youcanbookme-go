package util


type ProfileAvailableAccounts struct {
	Accountid      string `json:"accountId"`
				Email          string `json:"email"`
				Remoteaccounts []struct {
					Accountemail string `json:"accountEmail"`
					Calendars    []struct {
						Calendar ProfileCalendar `json:"calendar"`
						ID       string `json:"id"`
						Timezone struct {
						} `json:"timeZone"`
						Title     string `json:"title"`
						URL       string `json:"url"`
						Username  string `json:"userName"`
						Writeable bool   `json:"writeable"`
					} `json:"calendars"`
					ID                string `json:"id"`
					Lobbysupport      bool   `json:"lobbySupport"`
					Localaccountemail string `json:"localAccountEmail"`
					Type              string `json:"type"`
					Username          string `json:"username"`
				} `json:"remoteAccounts"`
}