# easyhttp

An extended easily mockable native like http client

# Basic Usage

```go
package main

import "github.com/mbirinci/easyhttp"

func main() {
	
	client := easyhttp.Client{ // you can use it like stdlib *http.Client
		&http.Client{}
	}
	
	resp, err := client.EasyGet("http://domain.tld/path")
	
	if err != nil {
		panic(err)
	}
	
	var body struct{Foo string} {Foo: "foo"}
	
	res.JSON(&body)
	
	fmt.Println("%v", body)
}

``` 

### Mock Easily

```go
package application

type HttpClient interface {
	Get(url string) (*easyhttp.Response, error)
}

type Application struct{
	client HttpClient
}

func NewApp(c HttpClient) *Application {
	return Application{c}
}

type Foo struct{
	Boo string
}

func (a *Application) Run() Foo {
	
	res, err := a.client.Get("http://app.server/path")
	
	if err != nil {
		panic(err)
	}
	
	var foo Foo
	
	err = res.JSON(&foo)
	
	if err != nil {
    		panic(err)
    }
	
	return foo
}

``` 

```go
package application_test

type mockHttpClient struct{}

func (*mockHttpClient) Get(url string) (*easyhttp.Response, error) {
	return &easyhttp.Response{RawBody: []byte(`"boo": "boo"`)}, nil
}

func TestApp(t *testing.T) {
	
	app := Application{HttpClient: mockHttpClient}
	
	f := app.Run()
	
	if f.Boo != "boo" {
		t.Fatalf("expected boo, but got %s", f.Boo)
	}
	
}

``` 
