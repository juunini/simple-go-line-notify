package notify

import (
	"net/url"
	"strings"
)

// SendImage : send line notify simple text with image
func SendImage(accessToken, message, imageURL string) (err error) {
	body := strings.NewReader(url.Values{
		"message":        []string{message},
		"imageThumbnail": []string{imageURL},
		"imageFullsize":  []string{imageURL},
	}.Encode())
	contentType := "application/x-www-form-urlencoded"

	err = sendToLineServer(body, accessToken, contentType)
	return
}
