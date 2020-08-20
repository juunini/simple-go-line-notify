package notify

import (
	"net/http"
	"net/url"
	"strings"
)

// SendImage : send line notify simple text with image
func SendImage(accessToken, message, imageURL string) (err error) {
	req, err := http.NewRequest(
		"POST",
		lineNotifyApiURL,
		strings.NewReader(url.Values{
			"message":        []string{message},
			"imageThumbnail": []string{imageURL},
			"imageFullsize":  []string{imageURL},
		}.Encode()),
	)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	err = sendToLineServer(req, accessToken)
	return
}
