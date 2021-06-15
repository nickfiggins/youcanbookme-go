package util

type Teammember struct {
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
}