package domain

type ConnectedEvent struct {
	ClientID            string `json:"clientId"`
	Timestamp           int64  `json:"timestamp"`
	EventType           string `json:"eventType"`
	SessionIdentifier   string `json:"sessionIdentifier"`
	PrincipalIdentifier string `json:"principalIdentifier"`
	IPAddress           string `json:"ipAddress"`
	VersionNumber       int64  `json:"versionNumber"`
}
