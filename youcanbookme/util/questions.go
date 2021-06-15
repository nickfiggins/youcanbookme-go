package util

type Question struct {
	After             string   `json:"after"`
	Before            string   `json:"before"`
	Code              string   `json:"code"`
	Options           []string `json:"options"`
	Required          bool     `json:"required"`
	Type              string   `json:"type"`
	Validation        string   `json:"validation"`
	Validationmessage string   `json:"validationMessage"`
}