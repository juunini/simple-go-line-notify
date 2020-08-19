package notify_test

import "github.com/juunini/simple-go-line-notify/notify"

func ExampleSendText() {
	bearer := "29jWoO****p70eK3AKA********ooHfusvDP6***ZmR"
	message := "hello, world!"

	if err := notify.SendText(bearer, message); err != nil {
		panic(err)
	}
}