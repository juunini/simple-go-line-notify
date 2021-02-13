package notify

import (
	"net/url"
	"strings"
)

// SendText : send line notify simple text
func SendText(accessToken, message string) (err error) {
	body := strings.NewReader(url.Values{
		"message": []string{message},
	}.Encode())
	contentType := "application/x-www-form-urlencoded"

	err = sendToLineServer(body, accessToken, contentType)
	return
}
