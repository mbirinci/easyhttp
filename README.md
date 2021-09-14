# easyhttp

[![codecov](https://codecov.io/gh/mbirinci/easyhttp/branch/master/graph/badge.svg?token=WGSYR5V8ZH)](https://codecov.io/gh/mbirinci/easyhttp)
[![Build Status](https://github.com/mbirinci/easyhttp/workflows/Test%20and%20Coverage/badge.svg?branch=master)](https://github.com/mbirinci/easyhttp/actions?query=branch%3Amaster)
[![GoDoc](https://pkg.go.dev/badge/github.com/mbirinci/easyhttp?status.svg)](https://pkg.go.dev/github.com/mbirinci/easyhttp?tab=doc)

An extended easily mockable native compatible http client

### Basic Usage

```go
package main

import "github.com/mbirinci/easyhttp"

func main() {
	
	client := easyhttp.Client {
		&http.Client{} // pass your underlying stdlib *http.Client
	}
	
	resp, err := client.EasyGet("http://foo.bar/", &easyhttp.Options{
		map[string]string{
			"If-Not-Now": "When",
		},
	})
	
	if err != nil {
		panic(err)
	}
	
	var data struct{ Foo string }
	
	resp.JSON(&data)
	
	// use the data
	fmt.Println("%v", data)
}

``` 

### Mock Easily

```go
package application

type HttpClient interface {
	EasyGet(url string, opts *easyhttp.Options) (*easyhttp.Response, error)
}

type Application struct {
	httpClient HttpClient
}

func NewApp(c HttpClient) *Application {
	return Application{ c }
}

type Foo struct{
	Bar string
}

func (app *Application) GetFoo() Foo {
	
	resp, err := app.httpClient.EasyGet("http://foo.bar", &easyhttp.Options{})
	
	if err != nil {
		panic(err)
	}
	
	var foo Foo
	
	err = resp.JSON(&foo)
	
	if err != nil {
	    panic(err)
	}
	
	return foo
}

``` 

```go
package application_test

type mockHttpClient struct{}

func (*mockHttpClient) EasyGet(url string) (*easyhttp.Response, error) {
	
	return &easyhttp.Response{RawBody: []byte(`{"bar": "bar"}`)}, nil
	
}

func TestApp(t *testing.T) {
	
	app := Application{HttpClient: mockHttpClient}
	
	foo := app.GetFoo()
	
	if foo.Bar != "bar" {
		t.Fatalf("expected bar, but got %s", foo.Bar)
	}
	
}

``` 
