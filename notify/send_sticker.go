package notify

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// SendSticker : send line notify simple text with sticker
//
// Sticker List
//
// https://devdocs.line.me/files/sticker_list.pdf
func SendSticker(accessToken, message string, stickerPackageId, stickerId int) (err error) {
	req, err := http.NewRequest(
		"POST",
		lineNotifyApiURL,
		strings.NewReader(url.Values{
			"message":          []string{message},
			"stickerPackageId": []string{strconv.Itoa(stickerPackageId)},
			"stickerId":        []string{strconv.Itoa(stickerId)},
		}.Encode()),
	)
	if err != nil {
		return
	}

	err = sendToLineServer(req, accessToken)
	return
}
