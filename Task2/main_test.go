
package main

import (
	
	"net/http"
	"net/http/httptest"
	"testing"
)

// for /articles/ route

func TestShowRoute(t *testing.T){

	testCases := []struct {
		name   string
	}{
		{name: "test1"},
		{ name: "test2"},
	}


	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			req, err := http.NewRequest("GET", "http://localhost:10000/articles/", nil)
			
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			res := httptest.NewRecorder()
			handler := http.HandlerFunc(articleFunction)
			handler.ServeHTTP(res, req)
			
			if status := res.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

		})
	}


}

// for /articles/:id route

func TestIdRoute(t *testing.T) {

	testCases := []struct {
		Id  string
		name   string
	}{
		{Id: "1", name: "test1"},
		{Id: "2", name: "test2"},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			req, err := http.NewRequest("GET", "http://localhost:10000/articles/"+ tc.Id, nil)
			
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			res := httptest.NewRecorder()
			handler := http.HandlerFunc(articleFunction)
			handler.ServeHTTP(res, req)
			
			if status := res.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

		})
	}
}

// for /articles/search?q=<key> route

func TestQueryRoute(t *testing.T) {

	testCases := []struct {
		query  string
		name   string
	}{
		{query: "Hello1", name: "test1"},
		{query: "Article Content2", name: "test2"},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			req, err := http.NewRequest("GET", "http://localhost:10000/articles/search?="+ tc.query, nil)
			
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			res := httptest.NewRecorder()
			handler := http.HandlerFunc(articleFunction)
			handler.ServeHTTP(res, req)
			
			if status := res.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

		})
	}
}







