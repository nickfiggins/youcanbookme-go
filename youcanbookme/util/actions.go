package util

import (
	"time"
)

type Action struct {
	Accountid                  string    `json:"accountId"`
		Anchor                     string    `json:"anchor"`
		Attachics                  bool      `json:"attachIcs"`
		Body                       string    `json:"body"`
		Bookingid                  string    `json:"bookingId"`
		Cc                         string    `json:"cc"`
		Created                    time.Time `json:"created"`
		Creditsused                int       `json:"creditsUsed"`
		Defaultasstring            string    `json:"defaultAsString"`
		Deletable                  bool      `json:"deletable"`
		Displayfiredatmediummedium string    `json:"displayFiredAtMediumMedium"`
		Displayfiredatmediumshort  string    `json:"displayFiredAtMediumShort"`
		Displayfiredatshortshort   string    `json:"displayFiredAtShortShort"`
		Displaylocale              struct {
		} `json:"displayLocale"`
		Displaytimezone struct {
		} `json:"displayTimeZone"`
		Displaytypelong  string `json:"displayTypeLong"`
		Displaytypeshort string `json:"displayTypeShort"`
		Editable         bool   `json:"editable"`
		Emailid          string `json:"emailId"`
		Failurecode      string `json:"failureCode"`
		Failuremessage   string `json:"failureMessage"`
		Fireat           struct {
		} `json:"fireAt"`
		Fireatforceweekday bool      `json:"fireAtForceWeekday"`
		Firedat            time.Time `json:"firedAt"`
		Fromaddress        string    `json:"fromAddress"`
		Fromname           string    `json:"fromName"`
		ID                 string    `json:"id"`
		Imageurl           string    `json:"imageUrl"`
		Locale             struct {
		} `json:"locale"`
		Offsetminutes   int    `json:"offsetMinutes"`
		Parentid        string `json:"parentId"`
		Profileid       string `json:"profileId"`
		Status          string `json:"status"`
		Stylingtemplate string `json:"stylingTemplate"`
		Subject         string `json:"subject"`
		Timezone        struct {
		} `json:"timeZone"`
		Timeline    string    `json:"timeline"`
		Title       string    `json:"title"`
		To          string    `json:"to"`
		Type        string    `json:"type"`
		Updated     time.Time `json:"updated"`
		Withinquota bool      `json:"withinQuota"`
		Ycbmbranded bool      `json:"ycbmBranded"`
}