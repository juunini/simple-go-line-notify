// Use Line Notify REST API
//
// See Also
//
// https://engineering.linecorp.com/en/blog/using-line-notify-to-send-messages-to-line-from-the-command-line/
package notify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// SendText : send line notify simple text
func SendText(Bearer, message string) (err error) {
	req, err := http.NewRequest("POST", "https://notify-api.line.me/api/notify", strings.NewReader(url.Values{"message": []string{message}}.Encode()))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+Bearer)

	client := &http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	res, err := client.Do(req)
	if err != nil {
		return
	}

	var body struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	if err = json.NewDecoder(res.Body).Decode(&body); err != nil {
		return
	}
	defer res.Body.Close()

	if body.Status != 200 {
		err = fmt.Errorf("%d: %s", body.Status, body.Message)
	}
	return
}
