package youcanbookme

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dghubble/sling"
	"github.com/nickfiggins/youcanbookme-go/youcanbookme/util"
)

type PremiumBooking struct {
	Acceptedat time.Time `json:"acceptedAt,omitempty"`
	Accountid  string    `json:"accountId,omitempty"`
	Actions    []util.Action `json:"actions,omitempty"`
	Answers    []util.Answer `json:"answers,omitempty"`
	Appointmenttypeids []string `json:"appointmentTypeIds,omitempty"`
	Appointmenttypes   []util.AppointmentType `json:"appointmentTypes,omitempty"`
	Appointmenttypesids  []string  `json:"appointmentTypesIds,omitempty"`
	Bookersecret         string    `json:"bookerSecret,omitempty"`
	Cancelreasonrequired bool      `json:"cancelReasonRequired,omitempty"`
	Cancellable          bool      `json:"cancellable,omitempty"`
	Cancellationcode     string    `json:"cancellationCode,omitempty"`
	Cancellationreason   string    `json:"cancellationReason,omitempty"`
	Cancelled            bool      `json:"cancelled,omitempty"`
	Cancelledat          time.Time `json:"cancelledAt,omitempty"`
	Cancelledby          string    `json:"cancelledBy,omitempty"`
	Createdat            time.Time `json:"createdAt,omitempty"`
	Currency             string    `json:"currency,omitempty"`
	Currencyfactor       int       `json:"currencyFactor,omitempty"`
	Databaseversion      util.DatabaseVersion `json:"databaseVersion,omitempty"`
	Displaybookerendsatmediumshort   string `json:"displayBookerEndsAtMediumShort,omitempty"`
	Displaybookerstartsatmediumshort string `json:"displayBookerStartsAtMediumShort,omitempty"`
	Displaycreatedatmedium           string `json:"displayCreatedAtMedium,omitempty"`
	Displaycreatedatshortshort       string `json:"displayCreatedAtShortShort,omitempty"`
	Displaydurationfull              string `json:"displayDurationFull,omitempty"`
	Displaydurationshort             string `json:"displayDurationShort,omitempty"`
	Displayendsatmediumshort         string `json:"displayEndsAtMediumShort,omitempty"`
	Displaylocale                    struct {
	} `json:"displayLocale,omitempty"`
	Displayprice                            string `json:"displayPrice,omitempty"`
	Displaystartsatfullshortanddurationfull string `json:"displayStartsAtFullShortAndDurationFull,omitempty"`
	Displaystartsatmediumshort              string `json:"displayStartsAtMediumShort,omitempty"`
	Displaytimezone                         struct {
	} `json:"displayTimeZone,omitempty"`
	Durationminutes int       `json:"durationMinutes,omitempty"`
	Endsat          time.Time `json:"endsAt,omitempty"`
	Endsatutc       time.Time `json:"endsAtUTC,omitempty"`
	Event           struct {
	} `json:"event,omitempty"`
	Expectedprice int       `json:"expectedPrice,omitempty"`
	ID            string    `json:"id,omitempty"`
	Initiatedat   time.Time `json:"initiatedAt,omitempty"`
	IP            string    `json:"ip,omitempty"`
	Linkfields    string    `json:"linkFields,omitempty"`
	Lobby         struct {
		Outboundurl string `json:"outboundUrl,omitempty"`
		Status      string `json:"status,omitempty"`
	} `json:"lobby,omitempty"`
	Lobbypassword string `json:"lobbyPassword,omitempty"`
	Locale        struct {
	} `json:"locale,omitempty"`
	Noshow        bool   `json:"noShow,omitempty"`
	Numberofslots int    `json:"numberOfSlots,omitempty"`
	Ownersecret   string `json:"ownerSecret,omitempty"`
	Paymentid     string `json:"paymentId,omitempty"`
	Preview       bool   `json:"preview,omitempty"`
	Price         int    `json:"price,omitempty"`
	Profile       util.BookingProfile `json:"profile,omitempty"`
	Profileaccesscode string    `json:"profileAccessCode,omitempty"`
	Profileid         string    `json:"profileId,omitempty"`
	Profilepassword   string    `json:"profilePassword,omitempty"`
	Recaptcharesponse string    `json:"recaptchaResponse,omitempty"`
	Ref               string    `json:"ref,omitempty"`
	Rejectedat        time.Time `json:"rejectedAt,omitempty"`
	Remoteid          string    `json:"remoteId,omitempty"`
	Rescheduledat     time.Time `json:"rescheduledAt,omitempty"`
	Rescheduledby     string    `json:"rescheduledBy,omitempty"`
	Reviewat          time.Time `json:"reviewAt,omitempty"`
	Serviceids        []string  `json:"serviceIds,omitempty"`
	Services          []struct {
		Description       string `json:"description,omitempty"`
		ID                string `json:"id,omitempty"`
		Name              string `json:"name,omitempty"`
		Numberofslots     int    `json:"numberOfSlots,omitempty"`
		Pic               string `json:"pic,omitempty"`
		Price             int    `json:"price,omitempty"`
		Slotlengthminutes int    `json:"slotLengthMinutes,omitempty"`
	} `json:"services,omitempty"`
	Startrescheduledfrom    time.Time `json:"startRescheduledFrom,omitempty"`
	Startrescheduledfromutc struct {
	} `json:"startRescheduledFromUTC,omitempty"`
	Startsat    time.Time `json:"startsAt,omitempty"`
	Startsatutc time.Time `json:"startsAtUTC,omitempty"`
	Stripetoken string    `json:"stripeToken,omitempty"`
	Teammember  util.Teammember `json:"teamMember,omitempty"`
	Teammemberid string `json:"teamMemberId,omitempty"`
	Tentative    string `json:"tentative,omitempty"`
	Timezone     struct {
	} `json:"timeZone,omitempty"`
	Timeline  string `json:"timeline,omitempty"`
	Title     string `json:"title,omitempty"`
	Units     int    `json:"units,omitempty"`
	Updatedat struct {
	} `json:"updatedAt,omitempty"`
}

type Booking struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Accountid string    `json:"accountId"`
	Profileid string    `json:"profileId"`
	Createdat time.Time `json:"createdAt"`
	Startsat  string    `json:"startsAt"`
	Endsat    string    `json:"endsAt"`
	Tentative string    `json:"tentative"`
	Timezone  string    `json:"timeZone"`
	Cancelled bool      `json:"cancelled"`
}


type BookingsService struct {
	sling *sling.Sling
	CurrentUser Account
}

func newBookingsService(sling *sling.Sling, user Account) *BookingsService {
	return &BookingsService{
		sling: sling.Path("bookings/"),
		CurrentUser: user,
	}
}

func (s *BookingsService) GetBookings() ([]Booking, *http.Response, error) {
	apiError := new(APIError)
	var bookings []Booking
	resp, err := s.sling.New().Receive(&bookings, apiError); if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	return bookings, resp, relevantError(err, *apiError)
}
