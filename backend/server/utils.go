package server

type errorMessage struct {
	Error string `json:"error"`
}

func newErrorMessage(m string) *errorMessage {
	return &errorMessage{
		Error: m,
	}
}
