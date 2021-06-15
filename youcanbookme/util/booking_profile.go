package util

import "time"

type BookingProfile struct {
	Accesscode                         string `json:"accessCode"`
	Accountcanaccessrestrictedfeatures bool   `json:"accountCanAccessRestrictedFeatures"`
	Accountemail                       string `json:"accountEmail"`
	Accountfamilyname                  string `json:"accountFamilyName"`
	Accountgivenname                   string `json:"accountGivenName"`
	Accountid                          string `json:"accountId"`
	Accountlocale                      struct {
	} `json:"accountLocale"`
	Accountmobile   string `json:"accountMobile"`
	Accounttimezone struct {
	} `json:"accountTimeZone"`
	Actions []Action `json:"actions"`
	Afterwards Afterwards `json:"afterwards"`
	Appointmenttypes ProfileAppointmentType `json:"appointmentTypes"`
	Availableaccounts []ProfileAvailableAccounts `json:"availableAccounts"`
	Calendars struct {
		Calendarids    []string `json:"calendarIds"`
		Targetcalendar struct {
		} `json:"targetCalendar"`
		Targetcalendarid string `json:"targetCalendarId"`
	} `json:"calendars"`
	Cancelorreschedule struct {
		Allowed                  bool   `json:"allowed"`
		Cancellationinstructions string `json:"cancellationInstructions"`
		Limitmessage             string `json:"limitMessage"`
		Limitminutes             int    `json:"limitMinutes"`
		Reasonrequired           bool   `json:"reasonRequired"`
		Showreasontextbox        bool   `json:"showReasonTextBox"`
	} `json:"cancelOrReschedule"`
	Captchaactive   bool      `json:"captchaActive"`
	Createdat       time.Time `json:"createdAt"`
	Createdby       string    `json:"createdBy"`
	Databaseversion struct {
	} `json:"databaseVersion"`
	Datespattern string `json:"datesPattern"`
	Description  string `json:"description"`
	Display      BookingProfileDisplay `json:"display"`
	ID     string `json:"id"`
	Locale struct {
	} `json:"locale"`
	Logo     string `json:"logo"`
	Parentid string `json:"parentId"`
	Password string `json:"password"`
	Payments struct {
		Active             bool   `json:"active"`
		Currency           string `json:"currency"`
		Currencyfactor     int    `json:"currencyFactor"`
		Partner            string `json:"partner"`
		Partnerdescription string `json:"partnerDescription"`
		Priceperslot       int    `json:"pricePerSlot"`
		Roundingprecision  int    `json:"roundingPrecision"`
	} `json:"payments"`
	Protectedbyaccesscode bool `json:"protectedByAccessCode"`
	Questions             []struct {
		After             string   `json:"after"`
		Before            string   `json:"before"`
		Code              string   `json:"code"`
		Options           []string `json:"options"`
		Required          bool     `json:"required"`
		Type              string   `json:"type"`
		Validation        string   `json:"validation"`
		Validationmessage string   `json:"validationMessage"`
	} `json:"questions"`
	Realtimetopic   string `json:"realtimeTopic"`
	Remotereminders []struct {
		ID            string `json:"id"`
		Offsetminutes int    `json:"offsetMinutes"`
		Type          string `json:"type"`
	} `json:"remoteReminders"`
	Requestedaction string `json:"requestedAction"`
	Services        struct {
		Active        bool   `json:"active"`
		Combinable    bool   `json:"combinable"`
		Description   string `json:"description"`
		Groupsactive  bool   `json:"groupsActive"`
		Groupsdivider string `json:"groupsDivider"`
		Items         []struct {
			Description       string `json:"description"`
			ID                string `json:"id"`
			Name              string `json:"name"`
			Numberofslots     int    `json:"numberOfSlots"`
			Pic               string `json:"pic"`
			Price             int    `json:"price"`
			Slotlengthminutes int    `json:"slotLengthMinutes"`
		} `json:"items"`
		Randomizable bool `json:"randomizable"`
	} `json:"services"`
	Status      string `json:"status"`
	Subdomain   string `json:"subdomain"`
	Teammembers struct {
		Active                            bool   `json:"active"`
		Allocationstrategy                string `json:"allocationStrategy"`
		Allowteammemberchangeonreschedule bool   `json:"allowTeamMemberChangeOnReschedule"`
		Assumenopreferenceoption          bool   `json:"assumeNoPreferenceOption"`
		Description                       string `json:"description"`
		Includenopreferenceoption         bool   `json:"includeNoPreferenceOption"`
		Items                             []Teammember `json:"items"`
		Randomorder bool `json:"randomOrder"`
	} `json:"teamMembers"`
	Tentative struct {
		Active                 bool `json:"active"`
		Autorejectactive       bool `json:"autoRejectActive"`
		Autorejectafterminutes int  `json:"autoRejectAfterMinutes"`
	} `json:"tentative"`
	Timezone struct {
	} `json:"timeZone"`
	Timezoneoverride bool `json:"timeZoneOverride"`
	Times          BookingTimes `json:"times"`
	Timespattern   string    `json:"timesPattern"`
	Title          string    `json:"title"`
	Updatedat      time.Time `json:"updatedAt"`
	Verifiedemails []string  `json:"verifiedEmails"`
	Vouchers       struct {
		Active                    bool     `json:"active"`
		Codealreadyusedmessage    string   `json:"codeAlreadyUsedMessage"`
		Voucheralreadyusedmessage string   `json:"voucherAlreadyUsedMessage"`
		Vouchers                  []string `json:"vouchers"`
	} `json:"vouchers"`
}