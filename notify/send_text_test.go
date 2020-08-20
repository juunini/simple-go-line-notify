package notify_test

import (
	"github.com/juunini/simple-go-line-notify/notify"
	"testing"
)

func TestSendTextInvalidAccessToken(t *testing.T) {
	accessToken := "29jWoO****p70eK3AKA********ooHfusvDP6***ZmR"
	message := "hello, world!"

	err := notify.SendText(accessToken, message)
	if !(err != nil && err.Error() == "401: Invalid access token") {
		t.Log(err)
		t.Fail()
	}
}

func ExampleSendText() {
	accessToken := "29jWoO****p70eK3AKA********ooHfusvDP6***ZmR"
	message := "hello, world!"

	if err := notify.SendText(accessToken, message); err != nil {
		panic(err)
	}
}
