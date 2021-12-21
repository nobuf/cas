package cas

import "fmt"

// RequestError format from https://apiv2-doc.twitcasting.tv/#errors
type RequestError struct {
	Content    TwitCastingError `json:"error"`
	StatusCode int              // unserializable
}

type TwitCastingError struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details"`
}

func (tce *RequestError) Error() string {
	return fmt.Sprintf("status code %v: %v(%v)", tce.StatusCode, tce.Content.Message, tce.Content.Code)
}
