package domain

type DisconnectedRequest struct {
	ClientID                  string `json:"clientId"`
	Timestamp                 int64  `json:"timestamp"`
	EventType                 string `json:"eventType"`
	SessionIdentifier         string `json:"sessionIdentifier"`
	PrincipalIdentifier       string `json:"principalIdentifier"`
	ClientInitiatedDisconnect bool   `json:"clientInitiatedDisconnect"`
	DisconnectReason          string `json:"disconnectReason"`
	VersionNumber             int64  `json:"versionNumber"`
}
