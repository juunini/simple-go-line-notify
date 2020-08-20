package notify_test

import (
	"github.com/juunini/simple-go-line-notify/notify"
	"io/ioutil"
	"testing"
)

func ExampleSendLocalImage() {
	accessToken := "29jWoO****p70eK3AKA********ooHfusvDP6***ZmR"
	message := "hello, world!"
	imagePath := "./sample.png"

	if err := notify.SendLocalImage(accessToken, message, imagePath); err != nil {
		panic(err)
	}
}

func TestSendLocalImage(t *testing.T) {
	accessToken, err := ioutil.ReadFile("test_token.txt")
	if err != nil {
		return
	}
	message := "local image test"
	imagePath := "./sample.jpg"

	if err := notify.SendLocalImage(string(accessToken), message, imagePath); err != nil {
		t.Error(err)
		t.Fail()
	}
}
