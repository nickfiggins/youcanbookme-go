package util

type BookingProfileDisplay struct {
	Bustmobile             bool   `json:"bustMobile"`
	Colorbusyslot          string `json:"colorBusySlot"`
	Colorfreeslot          string `json:"colorFreeSlot"`
	Colorheader            string `json:"colorHeader"`
	CSS                    string `json:"css"`
	Fontheader             string `json:"fontHeader"`
	Fontheadercolor        string `json:"fontHeaderColor"`
	Fontheadersize         int    `json:"fontHeaderSize"`
	Fontparagraph          string `json:"fontParagraph"`
	Fontparagraphcolor     string `json:"fontParagraphColor"`
	Fontparagraphsize      int    `json:"fontParagraphSize"`
	Fontsub                string `json:"fontSub"`
	Fontsubcolor           string `json:"fontSubColor"`
	Fontsubsize            int    `json:"fontSubSize"`
	Footer                 string `json:"footer"`
	Header                 string `json:"header"`
	Periodshownperpagedays int    `json:"periodShownPerPageDays"`
	Showjumpdate           bool   `json:"showJumpDate"`
	Showperiods            bool   `json:"showPeriods"`
	Showtimezone           bool   `json:"showTimeZone"`
	Showuncheckeddays      bool   `json:"showUncheckedDays"`
	Theme                  string `json:"theme"`
}