package notify

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const lineNotifyApiURL = "https://notify-api.line.me/api/notify"

func sendToLineServer(req *http.Request, accessToken string) (err error) {
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+accessToken)

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
