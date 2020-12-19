package helpers

const (
	ERROR   = "error"
	MESSAGE = "message"
)

type response struct {
	MessageType string      `json:"message_type" xml:"message_type"`
	Message     string      `json:"message" xml:"message"`
	Error       bool        `json:"error" xml:"error"`
	Data        interface{} `json:"data" xml:"data"`
}

func ResponseJSON(
	messageType, message string,
	err bool,
	data interface{},
) *response {
	return &response{
		MessageType: messageType,
		Message:     message,
		Error:       err,
		Data:        data,
	}
}
