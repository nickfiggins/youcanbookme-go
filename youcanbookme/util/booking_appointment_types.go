package util

type AppointmentType struct {
	Description       string `json:"description"`
	ID                string `json:"id"`
	Name              string `json:"name"`
	Numberofslots     int    `json:"numberOfSlots"`
	Pic               string `json:"pic"`
	Price             int    `json:"price"`
	Slotlengthminutes int    `json:"slotLengthMinutes"`
}