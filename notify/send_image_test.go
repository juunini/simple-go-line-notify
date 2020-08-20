package notify_test

import (
	"github.com/juunini/simple-go-line-notify/notify"
	"io/ioutil"
	"testing"
)

func TestSendImage(t *testing.T) {
	accessToken, _ := ioutil.ReadFile("test_token.txt")
	imageURL := "https://d.line-scdn.net/n/line_lp/img/ogimage.png"

	if err := notify.SendImage(string(accessToken), "a", imageURL); err != nil {
		t.Error(err)
		t.Fail()
	}
}

func ExampleSendImage() {
	accessToken := "29jWoO****p70eK3AKA********ooHfusvDP6***ZmR"
	message := "hello, world!"
	imageURL := "https://d.line-scdn.net/n/line_lp/img/ogimage.png"

	if err := notify.SendImage(accessToken, message, imageURL); err != nil {
		panic(err)
	}
}
