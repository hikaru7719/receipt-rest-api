package form

// HttpChallenge
type HttpChallenge struct {
	Type      string `json:"type"`
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
}
