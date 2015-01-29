package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var stripeWebhookUrl = "/stripe"

func stripeTest(body []byte) (*httptest.ResponseRecorder, error) {
	r, err := http.NewRequest("POST", stripeWebhookUrl, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	recorder := httptest.NewRecorder()

	st := &stripeMux{Controller: controller}
	st.ServeHTTP(recorder, r)

	return recorder, nil
}

func TestStripe(t *testing.T) {
	Convey("Given webhook from stripe", t, func() {
		Convey("When webhook isn't implemented", func() {
			body := []byte(`{"type":"random.webhook"}`)
			recorder, err := stripeTest(body)

			So(err, ShouldBeNil)

			Convey("Then it should return 200", func() {
				So(recorder.Code, ShouldEqual, 200)
			})
		})
	})
}
