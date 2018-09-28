package twilio

import (
	"context"
	"net/url"
	"testing"
	"time"
)

func TestGetNumberPage(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()
	data := url.Values{"PageSize": []string{"1000"}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	numbers, err := envClient.IncomingNumbers.GetPage(ctx, data)
	if err != nil {
		t.Fatal(err)
	}
	if len(numbers.IncomingPhoneNumbers) == 0 {
		t.Error("expected to get a list of phone numbers, got back 0")
	}
}

func TestBuyNumber(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()
	_, err := envClient.IncomingNumbers.BuyNumber("+1foobar")
	if err == nil {
		t.Fatal("expected to get an error, got nil")
	}
	rerr, ok := err.(*TwilioError)
	if !ok {
		t.Fatal("couldn't cast err to a TwilioError")
	}
	expected := "+1foobar is not a valid number"
	if rerr.Message != expected {
		t.Errorf("expected Title to be %s, got %s", expected, rerr.Message)
	}
	if rerr.Status != 400 {
		t.Errorf("expected StatusCode to be 400, got %d", rerr.Status)
	}
}
