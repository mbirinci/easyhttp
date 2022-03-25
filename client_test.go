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

  res, err := client.EasyGet(s.URL, &Options{Header: map[string]string{"foo": "bar"}})

  if err != nil {
    t.Fatalf("err: %v", err)
  }

  if res == nil {
    t.Fatalf("response should not nil")
  }

}

func TestClient_EasyGet_Fail(t *testing.T) {

  client := &Client{&http.Client{}} // use as stdlib *http.Client

  res, err := client.EasyGet("://foo", nil)

  if err == nil {
    t.Fatalf("should return error")
  }

  if res != nil {
    t.Fatalf("response should nil")
  }

  res, err = client.EasyGet("http://notfound.server/", nil)

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

  res, err = client.EasyGet(s.URL, nil)

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

  res, err := client.EasyGet(s.URL, nil)

  if err != nil {
    t.Fatalf("err: %v", err)
  }

  var body struct{ Foo string }

  err = res.JSON(&body)

  if err != nil {
    t.Fatalf("err: %v", err)
  }

}

func TestClient_EasyGet_RequestHook(t *testing.T) {

  host := "custom-host"

  s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    if r.Host != host {
      t.Fatalf("custom hook error")
    }

  }))

  client := &Client{&http.Client{}}

  res, err := client.EasyGet(s.URL, &Options{
    RequestHook: func(r *http.Request) {
      r.Host = host
    },
  })

  if err != nil {
    t.Fatalf("err: %v", err)
  }

  if res == nil {
    t.Fatalf("response should not nil")
  }

  defer s.Close()

}

func TestResponse_JSON_Fail(t *testing.T) {

  s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    w.WriteHeader(http.StatusInternalServerError)

  }))

  defer s.Close()

  client := &Client{&http.Client{}}

  res, err := client.EasyGet(s.URL, nil)

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

func TestResponse_Text(t *testing.T) {

  s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    _, err := w.Write([]byte(`{"foo":"foo"}`))

    if err != nil {
      t.Fatalf("err: %v", err)
    }

  }))

  defer s.Close()

  client := &Client{&http.Client{}}

  res, err := client.EasyGet(s.URL, nil)

  if err != nil {
    t.Fatalf("err: %v", err)
  }

  if res.Text() != `{"foo":"foo"}` {
    t.Fatalf("err: %v", err)
  }

}

func TestClient_EasyHead(t *testing.T) {

  s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    _, err := w.Write([]byte(`{"foo": "foo"}`))

    if err != nil {
      t.Fatalf("err: %v", err)
    }

  }))

  defer s.Close()

  client := &Client{&http.Client{}}

  res, err := client.EasyHead(s.URL, &Options{Header: map[string]string{"foo": "bar"}})

  if err != nil {
    t.Fatalf("err: %v", err)
  }

  if res == nil {
    t.Fatalf("response should not nil")
  }

}

func TestClient_EasyHead_Fail(t *testing.T) {

  client := &Client{&http.Client{}} // use as stdlib *http.Client

  res, err := client.EasyHead("://foo", nil)

  if err == nil {
    t.Fatalf("should return error")
  }

  if res != nil {
    t.Fatalf("response should nil")
  }

  res, err = client.EasyHead("http://notfound.server/", nil)

  if err == nil {
    t.Fatalf("should return error")
  }

  if res != nil {
    t.Fatalf("response should nil")
  }

}

func TestClient_EasyHead_RequestHook(t *testing.T) {

  host := "custom-host"

  s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    if r.Host != host {
      t.Fatalf("custom hook error")
    }

  }))

  client := &Client{&http.Client{}}

  res, err := client.EasyHead(s.URL, &Options{
    RequestHook: func(r *http.Request) {
      r.Host = host
    },
  })

  if err != nil {
    t.Fatalf("err: %v", err)
  }

  if res == nil {
    t.Fatalf("response should not nil")
  }

  defer s.Close()

}
