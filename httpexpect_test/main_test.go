package main

//
//import (
//	"fmt"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//
//	"github.com/gavv/httpexpect"
//)
//
//var server *httptest.Server
//
//func TestBetter(t *testing.T) {
//	server = httptest.NewServer(setupRouter())
//	defer server.Close()
//
//	// create a test engine form server
//	e := httpexpect.New(t, server.URL)
//
//	cases := []struct {
//		caseName string
//		param    string
//		repKey   string
//		repValue string
//	}{
//		{
//			caseName: "test value",
//			param:    "test",
//			repKey:   "value",
//			repValue: "test",
//		},
//		{
//			caseName: "test failed",
//			param:    "蔡徐坤",
//			repKey:   "status",
//			repValue: "唱跳rap和篮球",
//		},
//		{
//			caseName: "test no value",
//			param:    "蔡徐坤",
//			repKey:   "status",
//			repValue: "no value",
//		},
//	}
//
//	for _, c := range cases {
//		t.Run(c.caseName, func(t *testing.T) {
//			e.GET(fmt.Sprintf("/user/%s", c.param)).Expect().Status(http.StatusOK).
//				JSON().Object().
//				ContainsKey("user").ValueEqual("user", c.param).
//				ContainsKey(c.repKey).ValueEqual(c.repKey, c.repValue)
//		})
//	}
//}
//
//func TestGin(t *testing.T) {
//	server = httptest.NewServer(setupRouter())
//	defer server.Close()
//
//	// create a test engine form server
//	e := httpexpect.New(t, server.URL)
//
//	// post
//	type PostBody struct {
//		User  string `json:"user"`
//		Value string `json:"value"`
//	}
//	postBody := PostBody{User: "name1", Value: "value1"}
//	t.Run("test post", func(t *testing.T) {
//		e.POST("/admin").WithJSON(postBody).Expect().
//			Status(http.StatusOK)
//	})
//
//	// get
//	t.Run("test get success", func(t *testing.T) {
//		e.GET(fmt.Sprintf("/user/%s", postBody.User)).Expect().
//			Status(http.StatusOK).
//			JSON().Object().
//			ContainsKey("user").ValueEqual("user", postBody.User).
//			ContainsKey("value").ValueEqual("value", postBody.Value)
//	})
//
//	t.Run("test get no value", func(t *testing.T) {
//		name := "蔡徐坤"
//		e.GET(fmt.Sprintf("/user/%s", name)).Expect().
//			Status(http.StatusOK).
//			JSON().Object().
//			ContainsKey("user").ValueEqual("user", name).
//			ContainsKey("status").ValueEqual("status", "no value")
//	})
//
//}
