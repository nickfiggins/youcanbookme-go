package util

type Afterwards struct {
	Bookerevent struct {
		Description  string   `json:"description"`
		Location     string   `json:"location"`
		Participants []string `json:"participants"`
		Title        string   `json:"title"`
	} `json:"bookerEvent"`
	Bookingdescription string `json:"bookingDescription"`
	Bookinglocation    string `json:"bookingLocation"`
	Bookingtitle       string `json:"bookingTitle"`
	Isurl              bool   `json:"isUrl"`
	Ownerevent         struct {
		Description  string   `json:"description"`
		Location     string   `json:"location"`
		Participants []string `json:"participants"`
		Title        string   `json:"title"`
	} `json:"ownerEvent"`
	Scrolltotop              bool   `json:"scrollToTop"`
	Text                     string `json:"text"`
	Updatewholebrowserwindow bool   `json:"updateWholeBrowserWindow"`
	URL                      string `json:"url"`
}