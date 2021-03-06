package entities

type UssdSession struct {
	SessionID string `json:"sessionId"`
	PhoneNumber string `json:"phoneNumber"`
	Text string `json:"text"`
}
