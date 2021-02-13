package notify

import (
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
	body := strings.NewReader(url.Values{
		"message":          []string{message},
		"stickerPackageId": []string{strconv.Itoa(stickerPackageId)},
		"stickerId":        []string{strconv.Itoa(stickerId)},
	}.Encode())
	contentType := "application/x-www-form-urlencoded"

	err = sendToLineServer(body, accessToken, contentType)
	return
}
