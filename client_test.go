package easyhttp

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_EasyGet(t *testing.T) {

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, err := w.Write([]byte(`{"foo": "foo"}`))

		if err != nil {
			t.Fatalf("err: %v", err)
		}

	}))

	defer s.Close()

	client := &Client{&http.Client{}}

	res, err := client.EasyGet(s.URL)

	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if res == nil {
		t.Fatalf("response should not nil")
	}

}

func TestClient_EasyGet_Fail(t *testing.T) {

	client := &Client{&http.Client{}} // use as stdlib *http.Client

	res, err := client.EasyGet("://foo")

	if err == nil {
		t.Fatalf("should return error")
	}

	if res != nil {
		t.Fatalf("response should nil")
	}

	res, err = client.EasyGet("http://notfound.server/")

	if err == nil {
		t.Fatalf("should return error")
	}

	if res != nil {
		t.Fatalf("response should nil")
	}

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "1")
	}))

	defer s.Close()

	res, err = client.EasyGet(s.URL)

	if err == nil {
		t.Fatalf("should return error")
	}

	if res != nil {
		t.Fatalf("response should nil")
	}

}

func TestResponse_JSON(t *testing.T) {

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, err := w.Write([]byte(`{"foo": "foo"}`))

		if err != nil {
			t.Fatalf("err: %v", err)
		}

	}))

	defer s.Close()

	client := &Client{&http.Client{}}

	res, err := client.EasyGet(s.URL)

	if err != nil {
		t.Fatalf("err: %v", err)
	}

	var body struct{ Foo string }

	err = res.JSON(&body)

	if err != nil {
		t.Fatalf("err: %v", err)
	}

}

func TestResponse_JSON_Fail(t *testing.T) {

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusInternalServerError)

	}))

	defer s.Close()

	client := &Client{&http.Client{}}

	res, err := client.EasyGet(s.URL)

	if err != nil {
		t.Fatalf("err: %v", err)
	}

	var body struct{ Foo string }

	err = res.JSON(&body)

	if err == nil {
		t.Fatalf("should return error")
	}

	if body != struct{ Foo string }{Foo: ""} {
		t.Fatalf("body should have default value")
	}

}
