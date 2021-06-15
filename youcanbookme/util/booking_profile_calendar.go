package util

import "time"

type ProfileCalendar struct {
	Backgroundcolor string `json:"backgroundColor"`
	Changesrarely   bool   `json:"changesRarely"`
	Deleted         bool   `json:"deleted"`
	Description     string `json:"description"`
	Events          []struct {
		Backgroundcolor string `json:"backgroundColor"`
		Calendar        struct {
		} `json:"calendar"`
		Confidential      bool      `json:"confidential"`
		Confidentialonly  bool      `json:"confidentialOnly"`
		Createdat         time.Time `json:"createdAt"`
		Dateonly          bool      `json:"dateOnly"`
		Defaultsharelevel string    `json:"defaultShareLevel"`
		Description       string    `json:"description"`
		End               string    `json:"end"`
		Enddate           string    `json:"endDate"`
		Enddatetime       time.Time `json:"endDateTime"`
		Endtimezone       struct {
		} `json:"endTimeZone"`
		Endzoneddatetime time.Time `json:"endZonedDateTime"`
		Failed           bool      `json:"failed"`
		Failure          string    `json:"failure"`
		Foregroundcolor  string    `json:"foregroundColor"`
		ID               string    `json:"id"`
		Links            []struct {
			Href string `json:"href"`
			ID   string `json:"id"`
			Rel  string `json:"rel"`
			Type string `json:"type"`
		} `json:"links"`
		Location string `json:"location"`
		Metadata struct {
		} `json:"metadata"`
		Participants []struct {
			Email string `json:"email"`
			ID    string `json:"id"`
			Links []struct {
				Href string `json:"href"`
				ID   string `json:"id"`
				Rel  string `json:"rel"`
				Type string `json:"type"`
			} `json:"links"`
			Required bool   `json:"required"`
			Role     string `json:"role"`
			Status   string `json:"status"`
		} `json:"participants"`
		Private     bool     `json:"private"`
		Privateonly bool     `json:"privateOnly"`
		Public      bool     `json:"public"`
		Publiconly  bool     `json:"publicOnly"`
		Recurrences []string `json:"recurrences"`
		Reminders   []struct {
			Form  string `json:"form"`
			ID    string `json:"id"`
			Links []struct {
				Href string `json:"href"`
				ID   string `json:"id"`
				Rel  string `json:"rel"`
				Type string `json:"type"`
			} `json:"links"`
			Seconds int `json:"seconds"`
		} `json:"reminders"`
		Restricted     bool      `json:"restricted"`
		Restrictedonly bool      `json:"restrictedOnly"`
		Revision       string    `json:"revision"`
		Start          string    `json:"start"`
		Startdate      string    `json:"startDate"`
		Startdatetime  time.Time `json:"startDateTime"`
		Starttimezone  struct {
		} `json:"startTimeZone"`
		Startzoneddatetime time.Time `json:"startZonedDateTime"`
		Status             string    `json:"status"`
		Title              string    `json:"title"`
		Transparent        bool      `json:"transparent"`
		Virtualmeetingurl  string    `json:"virtualMeetingUrl"`
	} `json:"events"`
	Failed          bool   `json:"failed"`
	Failure         string `json:"failure"`
	Foregroundcolor string `json:"foregroundColor"`
	Freebusy        bool   `json:"freeBusy"`
	Freebusyonly    bool   `json:"freeBusyOnly"`
	Hidden          bool   `json:"hidden"`
	ID              string `json:"id"`
	Links           []struct {
		Href string `json:"href"`
		ID   string `json:"id"`
		Rel  string `json:"rel"`
		Type string `json:"type"`
	} `json:"links"`
	Permissionlevel string `json:"permissionLevel"`
	Primary         bool   `json:"primary"`
	Readable        bool   `json:"readable"`
	Readableonly    bool   `json:"readableOnly"`
	Remoteaccount   struct {
		Accesstoken          string    `json:"accessToken"`
		Accesstokenexpiresat time.Time `json:"accessTokenExpiresAt"`
		Account              struct {
			Accountintegrations struct {
				Accountid                   string    `json:"accountId"`
				Gmailaccesstoken            string    `json:"gmailAccessToken"`
				Gmailaccesstokenexpires     time.Time `json:"gmailAccessTokenExpires"`
				Gmailaddress                string    `json:"gmailAddress"`
				Gmailrefreshtoken           string    `json:"gmailRefreshToken"`
				Salesforceaccesstoken       string    `json:"salesforceAccessToken"`
				Salesforceemailaddress      string    `json:"salesforceEmailAddress"`
				Salesforceinstanceurl       string    `json:"salesforceInstanceUrl"`
				Salesforcerefreshtoken      string    `json:"salesforceRefreshToken"`
				Stripeconnectaccesstoken    string    `json:"stripeConnectAccessToken"`
				Stripeconnectpublishablekey string    `json:"stripeConnectPublishableKey"`
				Stripeconnectrefreshtoken   string    `json:"stripeConnectRefreshToken"`
				Stripeconnectuserid         string    `json:"stripeConnectUserId"`
				Zapierkey                   string    `json:"zapierKey"`
				Zoomaccesstoken             string    `json:"zoomAccessToken"`
				Zoomaccesstokenexpires      time.Time `json:"zoomAccessTokenExpires"`
				Zoomemailaddress            string    `json:"zoomEmailAddress"`
				Zoomrefreshtoken            string    `json:"zoomRefreshToken"`
			} `json:"accountIntegrations"`
			Accounttype string `json:"accountType"`
			Allocations []struct {
				Caligraphid       string `json:"caligraphId"`
				Email             string `json:"email"`
				Ghost             bool   `json:"ghost"`
				ID                string `json:"id"`
				Parentcaligraphid string `json:"parentCaligraphId"`
				Parentid          string `json:"parentId"`
				Quantityallocated int    `json:"quantityAllocated"`
			} `json:"allocations"`
			Billingnameandaddress struct {
				Address          string `json:"address"`
				City             string `json:"city"`
				Country          string `json:"country"`
				Familyname       string `json:"familyName"`
				Givenname        string `json:"givenName"`
				Organizationname string `json:"organizationName"`
				Postalcode       string `json:"postalCode"`
				Region           string `json:"region"`
			} `json:"billingNameAndAddress"`
			Blocked     bool   `json:"blocked"`
			Caligraphid string `json:"caligraphId"`
			Cards       []struct {
				Brand     string `json:"brand"`
				Country   string `json:"country"`
				Expdate   string `json:"expDate"`
				ID        string `json:"id"`
				Isdefault bool   `json:"isDefault"`
				Last4     string `json:"last4"`
				Name      string `json:"name"`
				Token     string `json:"token"`
			} `json:"cards"`
			Children []struct {
				Email    string `json:"email"`
				ID       string `json:"id"`
				Parentid string `json:"parentId"`
				Profiles []struct {
				} `json:"profiles"`
			} `json:"children"`
			Clientstate             string    `json:"clientState"`
			Contactmobile           string    `json:"contactMobile"`
			Createdat               time.Time `json:"createdAt"`
			Credit                  int       `json:"credit"`
			Crediteur               int       `json:"creditEUR"`
			Creditgbp               int       `json:"creditGBP"`
			Creditusd               int       `json:"creditUSD"`
			Currency                string    `json:"currency"`
			Debugmodeuntil          time.Time `json:"debugModeUntil"`
			Deletebookingsafterdays int       `json:"deleteBookingsAfterDays"`
			Deletedbookingscount    int       `json:"deletedBookingsCount"`
			Email                   string    `json:"email"`
			Familyname              string    `json:"familyName"`
			Feecredit               int       `json:"feeCredit"`
			Ghost                   bool      `json:"ghost"`
			Givenname               string    `json:"givenName"`
			Gmailaddress            string    `json:"gmailAddress"`
			Googleappsaccount       bool      `json:"googleAppsAccount"`
			Gridbranding            bool      `json:"gridBranding"`
			Hasbookings             bool      `json:"hasBookings"`
			ID                      string    `json:"id"`
			Indebugmode             bool      `json:"inDebugMode"`
			Killbillid              string    `json:"killBillId"`
			Lastwarningpushat       time.Time `json:"lastWarningPushAt"`
			Lastwarningpushlevel    int       `json:"lastWarningPushLevel"`
			Lifecycle               string    `json:"lifecycle"`
			Locale                  struct {
			} `json:"locale"`
			Loggedin struct {
			} `json:"loggedIn"`
			Marketingconsent            bool   `json:"marketingConsent"`
			Mfamethod                   string `json:"mfaMethod"`
			Needssynctopurchases        bool   `json:"needsSyncToPurchases"`
			Negotiateddiscountpermyriad int    `json:"negotiatedDiscountPermyriad"`
			Nextpurchase                struct {
				Account struct {
				} `json:"account"`
				Accountdiscount                      int       `json:"accountDiscount"`
				Accountid                            string    `json:"accountId"`
				Active                               bool      `json:"active"`
				Additionaldiscount                   int       `json:"additionalDiscount"`
				Additionaldiscountpermyriadrequested int       `json:"additionalDiscountPermyriadRequested"`
				Addresshibernate                     string    `json:"addressHibernate"`
				Afterdiscounts                       int       `json:"afterDiscounts"`
				Balance                              int       `json:"balance"`
				Beforediscounts                      int       `json:"beforeDiscounts"`
				Cardid                               string    `json:"cardId"`
				Cityhibernate                        string    `json:"cityHibernate"`
				Countryhibernate                     string    `json:"countryHibernate"`
				Createdat                            time.Time `json:"createdAt"`
				Currency                             string    `json:"currency"`
				Currencyfactor                       int       `json:"currencyFactor"`
				Date                                 string    `json:"date"`
				Displaylocale                        struct {
				} `json:"displayLocale"`
				Familynamehibernate string `json:"familyNameHibernate"`
				Givennamehibernate  string `json:"givenNameHibernate"`
				Gross               int    `json:"gross"`
				Hastax              bool   `json:"hasTax"`
				ID                  string `json:"id"`
				Liableforvat        bool   `json:"liableForVAT"`
				Localaccountid      string `json:"localAccountId"`
				Longperioddiscount  int    `json:"longPeriodDiscount"`
				Months              int    `json:"months"`
				Nameandaddress      struct {
					Address          string `json:"address"`
					City             string `json:"city"`
					Country          string `json:"country"`
					Familyname       string `json:"familyName"`
					Givenname        string `json:"givenName"`
					Organizationname string `json:"organizationName"`
					Postalcode       string `json:"postalCode"`
					Region           string `json:"region"`
				} `json:"nameAndAddress"`
				Needspayment              bool      `json:"needsPayment"`
				Net                       int       `json:"net"`
				Number                    string    `json:"number"`
				Organizationnamehibernate string    `json:"organizationNameHibernate"`
				Payment                   int       `json:"payment"`
				Paymentintentid           string    `json:"paymentIntentId"`
				Postalcodehibernate       string    `json:"postalCodeHibernate"`
				Proratapermyriadused      int       `json:"proRataPermyriadUsed"`
				Product                   string    `json:"product"`
				Provisionend              time.Time `json:"provisionEnd"`
				Provisionstart            time.Time `json:"provisionStart"`
				Quantity                  int       `json:"quantity"`
				Refundcode                string    `json:"refundCode"`
				Regionhibernate           string    `json:"regionHibernate"`
				Requestedaction           string    `json:"requestedAction"`
				Setupmode                 string    `json:"setUpMode"`
				Shortperioddiscount       int       `json:"shortPeriodDiscount"`
				Status                    string    `json:"status"`
				Tax                       int       `json:"tax"`
				Taxnumber                 string    `json:"taxNumber"`
				Totaldiscount             int       `json:"totalDiscount"`
				Transactions              []struct {
					Account     string    `json:"account"`
					Accountid   string    `json:"accountId"`
					Createdat   time.Time `json:"createdAt"`
					Description string    `json:"description"`
					Gross       int       `json:"gross"`
					ID          string    `json:"id"`
					Linked      struct {
					} `json:"linked"`
					Localaccountid string `json:"localAccountId"`
					Net            int    `json:"net"`
					Purchase       struct {
					} `json:"purchase"`
					Purchaseid string    `json:"purchaseId"`
					Remoteid   string    `json:"remoteId"`
					Tax        int       `json:"tax"`
					Type       string    `json:"type"`
					Updatedat  time.Time `json:"updatedAt"`
				} `json:"transactions"`
				Transactionssorted []struct {
					Account     string    `json:"account"`
					Accountid   string    `json:"accountId"`
					Createdat   time.Time `json:"createdAt"`
					Description string    `json:"description"`
					Gross       int       `json:"gross"`
					ID          string    `json:"id"`
					Linked      struct {
					} `json:"linked"`
					Localaccountid string `json:"localAccountId"`
					Net            int    `json:"net"`
					Purchase       struct {
					} `json:"purchase"`
					Purchaseid string    `json:"purchaseId"`
					Remoteid   string    `json:"remoteId"`
					Tax        int       `json:"tax"`
					Type       string    `json:"type"`
					Updatedat  time.Time `json:"updatedAt"`
				} `json:"transactionsSorted"`
				Type           string    `json:"type"`
				Unitprice      int       `json:"unitPrice"`
				Updatedat      time.Time `json:"updatedAt"`
				Volumediscount int       `json:"volumeDiscount"`
			} `json:"nextPurchase"`
			Organisation string `json:"organisation"`
			Parent       struct {
				Email           string `json:"email"`
				ID              string `json:"id"`
				Quantitypaidfor int    `json:"quantityPaidFor"`
			} `json:"parent"`
			Parentemail      string `json:"parentEmail"`
			Parentforbilling struct {
				Email           string `json:"email"`
				ID              string `json:"id"`
				Quantitypaidfor int    `json:"quantityPaidFor"`
			} `json:"parentForBilling"`
			Parentid              string `json:"parentId"`
			Parentidforbilling    string `json:"parentIdForBilling"`
			Parentquantitypaidfor int    `json:"parentQuantityPaidFor"`
			Pastdue               bool   `json:"pastDue"`
			Pauseplanmonths       int    `json:"pausePlanMonths"`
			Pausequantity         int    `json:"pauseQuantity"`
			Pauseremainingseconds int    `json:"pauseRemainingSeconds"`
			Paused                bool   `json:"paused"`
			Permissionsin         []struct {
				Createdat time.Time `json:"createdAt"`
				Creator   struct {
				} `json:"creator"`
				Databaseversion struct {
				} `json:"databaseVersion"`
				Expiresat   time.Time `json:"expiresAt"`
				Fromaccount struct {
				} `json:"fromAccount"`
				Fromlocalaccount struct {
				} `json:"fromLocalAccount"`
				ID             string `json:"id"`
				Resolvedstatus string `json:"resolvedStatus"`
				Resource       string `json:"resource"`
				Status         string `json:"status"`
				Toaccount      struct {
				} `json:"toAccount"`
				Tolocalaccount struct {
				} `json:"toLocalAccount"`
				Type      string    `json:"type"`
				Updatedat time.Time `json:"updatedAt"`
			} `json:"permissionsIn"`
			Permissionsout []struct {
				Createdat time.Time `json:"createdAt"`
				Creator   struct {
				} `json:"creator"`
				Databaseversion struct {
				} `json:"databaseVersion"`
				Expiresat   time.Time `json:"expiresAt"`
				Fromaccount struct {
				} `json:"fromAccount"`
				Fromlocalaccount struct {
				} `json:"fromLocalAccount"`
				ID             string `json:"id"`
				Resolvedstatus string `json:"resolvedStatus"`
				Resource       string `json:"resource"`
				Status         string `json:"status"`
				Toaccount      struct {
				} `json:"toAccount"`
				Tolocalaccount struct {
				} `json:"toLocalAccount"`
				Type      string    `json:"type"`
				Updatedat time.Time `json:"updatedAt"`
			} `json:"permissionsOut"`
			Plan          string    `json:"plan"`
			Planexpiresat time.Time `json:"planExpiresAt"`
			Planmonths    int       `json:"planMonths"`
			Planmonthswas int       `json:"planMonthsWas"`
			Profiles      []struct {
			} `json:"profiles"`
			Purchases []struct {
				Account struct {
				} `json:"account"`
				Accountdiscount                      int       `json:"accountDiscount"`
				Accountid                            string    `json:"accountId"`
				Active                               bool      `json:"active"`
				Additionaldiscount                   int       `json:"additionalDiscount"`
				Additionaldiscountpermyriadrequested int       `json:"additionalDiscountPermyriadRequested"`
				Addresshibernate                     string    `json:"addressHibernate"`
				Afterdiscounts                       int       `json:"afterDiscounts"`
				Balance                              int       `json:"balance"`
				Beforediscounts                      int       `json:"beforeDiscounts"`
				Cardid                               string    `json:"cardId"`
				Cityhibernate                        string    `json:"cityHibernate"`
				Countryhibernate                     string    `json:"countryHibernate"`
				Createdat                            time.Time `json:"createdAt"`
				Currency                             string    `json:"currency"`
				Currencyfactor                       int       `json:"currencyFactor"`
				Date                                 string    `json:"date"`
				Displaylocale                        struct {
				} `json:"displayLocale"`
				Familynamehibernate string `json:"familyNameHibernate"`
				Givennamehibernate  string `json:"givenNameHibernate"`
				Gross               int    `json:"gross"`
				Hastax              bool   `json:"hasTax"`
				ID                  string `json:"id"`
				Liableforvat        bool   `json:"liableForVAT"`
				Localaccountid      string `json:"localAccountId"`
				Longperioddiscount  int    `json:"longPeriodDiscount"`
				Months              int    `json:"months"`
				Nameandaddress      struct {
					Address          string `json:"address"`
					City             string `json:"city"`
					Country          string `json:"country"`
					Familyname       string `json:"familyName"`
					Givenname        string `json:"givenName"`
					Organizationname string `json:"organizationName"`
					Postalcode       string `json:"postalCode"`
					Region           string `json:"region"`
				} `json:"nameAndAddress"`
				Needspayment              bool      `json:"needsPayment"`
				Net                       int       `json:"net"`
				Number                    string    `json:"number"`
				Organizationnamehibernate string    `json:"organizationNameHibernate"`
				Payment                   int       `json:"payment"`
				Paymentintentid           string    `json:"paymentIntentId"`
				Postalcodehibernate       string    `json:"postalCodeHibernate"`
				Proratapermyriadused      int       `json:"proRataPermyriadUsed"`
				Product                   string    `json:"product"`
				Provisionend              time.Time `json:"provisionEnd"`
				Provisionstart            time.Time `json:"provisionStart"`
				Quantity                  int       `json:"quantity"`
				Refundcode                string    `json:"refundCode"`
				Regionhibernate           string    `json:"regionHibernate"`
				Requestedaction           string    `json:"requestedAction"`
				Setupmode                 string    `json:"setUpMode"`
				Shortperioddiscount       int       `json:"shortPeriodDiscount"`
				Status                    string    `json:"status"`
				Tax                       int       `json:"tax"`
				Taxnumber                 string    `json:"taxNumber"`
				Totaldiscount             int       `json:"totalDiscount"`
				Transactions              []struct {
					Account     string    `json:"account"`
					Accountid   string    `json:"accountId"`
					Createdat   time.Time `json:"createdAt"`
					Description string    `json:"description"`
					Gross       int       `json:"gross"`
					ID          string    `json:"id"`
					Linked      struct {
					} `json:"linked"`
					Localaccountid string `json:"localAccountId"`
					Net            int    `json:"net"`
					Purchase       struct {
					} `json:"purchase"`
					Purchaseid string    `json:"purchaseId"`
					Remoteid   string    `json:"remoteId"`
					Tax        int       `json:"tax"`
					Type       string    `json:"type"`
					Updatedat  time.Time `json:"updatedAt"`
				} `json:"transactions"`
				Transactionssorted []struct {
					Account     string    `json:"account"`
					Accountid   string    `json:"accountId"`
					Createdat   time.Time `json:"createdAt"`
					Description string    `json:"description"`
					Gross       int       `json:"gross"`
					ID          string    `json:"id"`
					Linked      struct {
					} `json:"linked"`
					Localaccountid string `json:"localAccountId"`
					Net            int    `json:"net"`
					Purchase       struct {
					} `json:"purchase"`
					Purchaseid string    `json:"purchaseId"`
					Remoteid   string    `json:"remoteId"`
					Tax        int       `json:"tax"`
					Type       string    `json:"type"`
					Updatedat  time.Time `json:"updatedAt"`
				} `json:"transactionsSorted"`
				Type           string    `json:"type"`
				Unitprice      int       `json:"unitPrice"`
				Updatedat      time.Time `json:"updatedAt"`
				Volumediscount int       `json:"volumeDiscount"`
			} `json:"purchases"`
			Quantityallocated int    `json:"quantityAllocated"`
			Quantityforfree   int    `json:"quantityForFree"`
			Quantityfreetrial int    `json:"quantityFreeTrial"`
			Quantitypaidfor   int    `json:"quantityPaidFor"`
			Realtimetopic     string `json:"realtimeTopic"`
			Reducedpolicing   bool   `json:"reducedPolicing"`
			Remoteaccounts    []struct {
			} `json:"remoteAccounts"`
			Requestedaction        string    `json:"requestedAction"`
			Reviewat               time.Time `json:"reviewAt"`
			Salesforceemailaddress string    `json:"salesforceEmailAddress"`
			Sector                 string    `json:"sector"`
			Security               struct {
				Accountid             string    `json:"accountId"`
				Apikey                string    `json:"apiKey"`
				Apikeyactiveat        string    `json:"apiKeyActiveAt"`
				Caligraphid           string    `json:"caligraphId"`
				Email                 string    `json:"email"`
				Ghost                 bool      `json:"ghost"`
				Mfamethod             string    `json:"mfaMethod"`
				Onetimetoken          string    `json:"oneTimeToken"`
				Onetimetokenexpiresat time.Time `json:"oneTimeTokenExpiresAt"`
				Password              string    `json:"password"`
				Passwordhash          string    `json:"passwordHash"`
				Sessiontoken          string    `json:"sessionToken"`
				Sessiontokenexpiresat time.Time `json:"sessionTokenExpiresAt"`
				Totpsecret            string    `json:"totpSecret"`
				Totpsecretverifiedat  time.Time `json:"totpSecretVerifiedAt"`
			} `json:"security"`
			Smscredits             int    `json:"smsCredits"`
			Source                 string `json:"source"`
			Stripeid               string `json:"stripeId"`
			Stripeuserid           string `json:"stripeUserId"`
			Synctopurchaseschanges struct {
				Hasvalue        bool      `json:"hasValue"`
				Plan            string    `json:"plan"`
				Planexpiresat   time.Time `json:"planExpiresAt"`
				Quantitypaidfor int       `json:"quantityPaidFor"`
				Trialendsat     time.Time `json:"trialEndsAt"`
			} `json:"syncToPurchasesChanges"`
			Tags      string    `json:"tags"`
			Taxnumber string    `json:"taxNumber"`
			Terms     bool      `json:"terms"`
			Termsdate time.Time `json:"termsDate"`
			Timezone  struct {
			} `json:"timeZone"`
			Trialendsat time.Time `json:"trialEndsAt"`
			Type        string    `json:"type"`
			Updatedat   time.Time `json:"updatedAt"`
			Uploadcode  string    `json:"uploadCode"`
			Warnings    []struct {
				Code       string `json:"code"`
				Fixurl     string `json:"fixUrl"`
				Fixurltext string `json:"fixUrlText"`
				Frequency  int    `json:"frequency"`
				Level      int    `json:"level"`
				Longtext   string `json:"longText"`
				Mediumtext string `json:"mediumText"`
				Push       bool   `json:"push"`
				Shorttext  string `json:"shortText"`
			} `json:"warnings"`
			Welcomed         bool   `json:"welcomed"`
			Xeroid           string `json:"xeroId"`
			Zapierkey        string `json:"zapierKey"`
			Zoomemailaddress string `json:"zoomEmailAddress"`
		} `json:"account"`
		Accountemail string `json:"accountEmail"`
		Accountid    string `json:"accountId"`
		Calendarhome string `json:"calendarHome"`
		Calendars    []struct {
		} `json:"calendars"`
		Databaseversion struct {
		} `json:"databaseVersion"`
		Failed  bool   `json:"failed"`
		Failure string `json:"failure"`
		Host    string `json:"host"`
		ID      string `json:"id"`
		Links   []struct {
			Href string `json:"href"`
			ID   string `json:"id"`
			Rel  string `json:"rel"`
			Type string `json:"type"`
		} `json:"links"`
		Localaccountemail string `json:"localAccountEmail"`
		Localaccountid    string `json:"localAccountId"`
		Password          string `json:"password"`
		Principal         string `json:"principal"`
		Provider          struct {
			Flushcache    bool `json:"flushCache"`
			Remoteaccount struct {
			} `json:"remoteAccount"`
		} `json:"provider"`
		Refreshtoken               string `json:"refreshToken"`
		Requestedaction            string `json:"requestedAction"`
		Revision                   string `json:"revision"`
		Statuscanceleddeletesevent bool   `json:"statusCanceledDeletesEvent"`
		Type                       string `json:"type"`
		Upcalendarsrecordfailure   bool   `json:"upCalendarsRecordFailure"`
		Username                   string `json:"username"`
	} `json:"remoteAccount"`
	Revision       string `json:"revision"`
	Systemuniqueid string `json:"systemUniqueId"`
	Timezone       struct {
	} `json:"timeZone"`
	Title         string `json:"title"`
	Username      string `json:"userName"`
	Writeable     bool   `json:"writeable"`
	Writeableonly bool   `json:"writeableOnly"`
}