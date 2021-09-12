# easyhttp

[![codecov](https://codecov.io/gh/mbirinci/easyhttp/branch/master/graph/badge.svg?token=WGSYR5V8ZH)](https://codecov.io/gh/mbirinci/easyhttp)
[![Build Status](https://github.com/mbirinci/easyhttp/workflows/Test%20and%20Coverage/badge.svg?branch=master)](https://github.com/mbirinci/easyhttp/actions?query=branch%3Amaster)

An extended easily mockable native like http client

### Basic Usage

```go
package main

import "github.com/mbirinci/easyhttp"

func main() {
	
	client := easyhttp.Client{
		&http.Client{} // pass your underlying stdlib *http.Client
	}
	
	resp, err := client.EasyGet("http://foo.bar/path")
	
	if err != nil {
		panic(err)
	}
	
	var data struct{Foo string} {Foo: "bar"}
	
	resp.JSON(&data)
	
	// use the data
	fmt.Println("%v", data)
}

``` 

### Mock Easily

```go
package application

type HttpClient interface {
	EasyGet(url string) (*easyhttp.Response, error)
}

type Application struct{
	httpClient HttpClient
}

func NewApp(c HttpClient) *Application {
	return Application{c}
}

type Foo struct{
	Bar string
}

func (app *Application) Run() Foo {
	
	resp, err := app.httpClient.EasyGet("http://foo.bar/path")
	
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
	
	return &easyhttp.Response{RawBody: []byte(`"bar": "bar"`)}, nil
	
}

func TestApp(t *testing.T) {
	
	app := Application{HttpClient: mockHttpClient}
	
	foo := app.Run()
	
	if foo.Bar != "bar" {
		t.Fatalf("expected bar, but got %s", foo.Bar)
	}
	
}

``` 
