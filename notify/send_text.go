package notify

import (
	"net/http"
	"net/url"
	"strings"
)

// SendText : send line notify simple text
func SendText(accessToken, message string) (err error) {
	req, err := http.NewRequest(
		"POST",
		lineNotifyApiURL,
		strings.NewReader(url.Values{
			"message": []string{message},
		}.Encode()),
	)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	err = sendToLineServer(req, accessToken)
	return
}
