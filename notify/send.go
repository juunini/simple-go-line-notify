package notify

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func Send(Bearer, message string) (err error) {
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
		err = errors.New(fmt.Sprintf("%d: %s", body.Status, body.Message))
	}
	return
}
