package removetokens

const (
	TypeRemoveTokenByClientID           = "token:remove-by-client-id"
	TypeRemoveTokenBySubject            = "token:remove-by-subject"
	TypeRemoveTokenByClientIDAndSubject = "token:remove-by-client-id-and-subject"
)

type (
	TokenRemoveByClientID struct {
		ClientID string `json:"client_id"`
	}
	TokenRemoveBySubject struct {
		Subject string `json:"subject"`
	}
	TokenRemoveByClientIDAndSubject struct {
		TokenRemoveByClientID
		TokenRemoveBySubject
	}
)
