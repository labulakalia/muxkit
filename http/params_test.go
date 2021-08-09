package http

import (
	"net/http"
	"net/url"
	"testing"
)

func TestParse(t *testing.T) {
	type Hello struct {
		HelloId uint64 `json:"hello_id"`
		Name    string `json:"name"`
		IsTrue  bool   `json:"is_true"`
	}
	hello := Hello{}
	nurl := &url.Values{}
	nurl.Add("hello_id", "12")
	nurl.Add("name", "test")
	nurl.Add("is_true", "true")

	req, err := http.NewRequest(http.MethodGet, "/test?"+nurl.Encode(), nil)
	if err != nil {
		t.Fatal(err)
	}
	err = ParseParams(req, &hello)
	if err != nil {
		t.Fatal(err)
	}
	if hello.HelloId != 12 {
		t.Fatal("hello id is not 12")
	}
	if hello.Name != "test" {
		t.Fatal("name is not test")
	}
	if hello.IsTrue != true {
		t.Fatal("IsTrue is not true")
	}
	t.Log(hello)

}
