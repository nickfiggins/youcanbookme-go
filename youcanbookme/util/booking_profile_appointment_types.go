package util

type ProfileAppointmentType struct {
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
}