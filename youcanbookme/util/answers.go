package util

type Answer struct {
	Code     string `json:"code"`
	Question Question `json:"question"`
	String string `json:"string"`
}