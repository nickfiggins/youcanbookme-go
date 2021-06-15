package util

import "time"

type DatabaseVersion struct {
	Acceptedat time.Time `json:"acceptedAt"`
		Accountid  string    `json:"accountId"`
		Actions    []Action `json:"actions"`
		Answers []Answer `json:"answers"`
		Appointmenttypeids []string `json:"appointmentTypeIds"`
		Appointmenttypes   []AppointmentType `json:"appointmentTypes"`
		Appointmenttypesids  []string  `json:"appointmentTypesIds"`
		Bookersecret         string    `json:"bookerSecret"`
		Cancelreasonrequired bool      `json:"cancelReasonRequired"`
		Cancellable          bool      `json:"cancellable"`
		Cancellationcode     string    `json:"cancellationCode"`
		Cancellationreason   string    `json:"cancellationReason"`
		Cancelled            bool      `json:"cancelled"`
		Cancelledat          time.Time `json:"cancelledAt"`
		Cancelledby          string    `json:"cancelledBy"`
		Createdat            time.Time `json:"createdAt"`
		Currency             string    `json:"currency"`
		Currencyfactor       int       `json:"currencyFactor"`
		Databaseversion      struct {
		} `json:"databaseVersion"`
		Displaybookerendsatmediumshort   string `json:"displayBookerEndsAtMediumShort"`
		Displaybookerstartsatmediumshort string `json:"displayBookerStartsAtMediumShort"`
		Displaycreatedatmedium           string `json:"displayCreatedAtMedium"`
		Displaycreatedatshortshort       string `json:"displayCreatedAtShortShort"`
		Displaydurationfull              string `json:"displayDurationFull"`
		Displaydurationshort             string `json:"displayDurationShort"`
		Displayendsatmediumshort         string `json:"displayEndsAtMediumShort"`
		Displaylocale                    struct {
		} `json:"displayLocale"`
		Displayprice                            string `json:"displayPrice"`
		Displaystartsatfullshortanddurationfull string `json:"displayStartsAtFullShortAndDurationFull"`
		Displaystartsatmediumshort              string `json:"displayStartsAtMediumShort"`
		Displaytimezone                         struct {
		} `json:"displayTimeZone"`
		Durationminutes int       `json:"durationMinutes"`
		Endsat          time.Time `json:"endsAt"`
		Endsatutc       time.Time `json:"endsAtUTC"`
		Event           struct {
		} `json:"event"`
		Expectedprice int       `json:"expectedPrice"`
		ID            string    `json:"id"`
		Initiatedat   time.Time `json:"initiatedAt"`
		IP            string    `json:"ip"`
		Linkfields    string    `json:"linkFields"`
		Lobby         struct {
			Outboundurl string `json:"outboundUrl"`
			Status      string `json:"status"`
		} `json:"lobby"`
		Lobbypassword string `json:"lobbyPassword"`
		Locale        struct {
		} `json:"locale"`
		Noshow        bool   `json:"noShow"`
		Numberofslots int    `json:"numberOfSlots"`
		Ownersecret   string `json:"ownerSecret"`
		Paymentid     string `json:"paymentId"`
		Preview       bool   `json:"preview"`
		Price         int    `json:"price"`
		Profile       BookingProfile `json:"profile"`
		Profileaccesscode string    `json:"profileAccessCode"`
		Profileid         string    `json:"profileId"`
		Profilepassword   string    `json:"profilePassword"`
		Recaptcharesponse string    `json:"recaptchaResponse"`
		Ref               string    `json:"ref"`
		Rejectedat        time.Time `json:"rejectedAt"`
		Remoteid          string    `json:"remoteId"`
		Rescheduledat     time.Time `json:"rescheduledAt"`
		Rescheduledby     string    `json:"rescheduledBy"`
		Reviewat          time.Time `json:"reviewAt"`
		Serviceids        []string  `json:"serviceIds"`
		Services          []struct {
			Description       string `json:"description"`
			ID                string `json:"id"`
			Name              string `json:"name"`
			Numberofslots     int    `json:"numberOfSlots"`
			Pic               string `json:"pic"`
			Price             int    `json:"price"`
			Slotlengthminutes int    `json:"slotLengthMinutes"`
		} `json:"services"`
		Startrescheduledfrom    time.Time `json:"startRescheduledFrom"`
		Startrescheduledfromutc struct {
		} `json:"startRescheduledFromUTC"`
		Startsat    time.Time `json:"startsAt"`
		Startsatutc time.Time `json:"startsAtUTC"`
		Stripetoken string    `json:"stripeToken"`
		Teammember  struct {
			Account        string `json:"account"`
			Calendarid     string `json:"calendarId"`
			Description    string `json:"description"`
			Email          string `json:"email"`
			ID             string `json:"id"`
			Name           string `json:"name"`
			Pic            string `json:"pic"`
			Targetcalendar struct {
			} `json:"targetCalendar"`
			Targetcalendartimezone struct {
			} `json:"targetCalendarTimeZone"`
			Targetcalendarwriteable bool `json:"targetCalendarWriteable"`
		} `json:"teamMember"`
		Teammemberid string `json:"teamMemberId"`
		Tentative    string `json:"tentative"`
		Timezone     struct {
		} `json:"timeZone"`
		Timeline  string `json:"timeline"`
		Title     string `json:"title"`
		Units     int    `json:"units"`
		Updatedat struct {
		} `json:"updatedAt"`
}