package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestBadHTTPHeader(t *testing.T) {
	//mux := http.NewServeMux()
	//mux.Handle("/", http.FileServer(http.Dir("."))

	//srv := &http.Server{
	//	Addr: ":8080",
	//	Handler: mux
	//}

	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Errorf("Oopsies bad request bad test case")
	}

	//t.Errorf(req.Header.Get("Authorization"))

	_, err2 := GetAPIKey(req.Header)

	if err2 != nil {
		fmt.Printf("Get API Key failed due to missing Authorization in Header!  Expected! Err: %s\n", err2)
	} else {
		t.Errorf("Should produce error 'Missing Authorization in Header'")
	}

	req.Header.Set("Authorization", "Bearer YTIwfAKeNGYtODY4YS00MzM5LTkzNGYtNGRhMmQ3ODhkMGFhOjNuYU8xMElBMmFnY3ZHKzlJOVRHRVE9PQ==")
	//t.Errorf(req.Header.Get("Authorization"))

	_, err3 := GetAPIKey(req.Header)

	if err3 != nil {
		fmt.Printf("Authorization not an ApiKey!  Expected! Err: %s\n", err3)
	} else {
		t.Errorf("Should produce error 'malformed authorization header'")
	}

	req.Header.Set("Authorization", "ApiKey YTIwfAKeNGYtODY4YS00MzM5LTkzNGYtNGRhMmQ3ODhkMGFhOjNuYU8xMElBMmFnY3ZHKzlJOVRHRVE9PQ==")

	apiKey, err4 := GetAPIKey(req.Header)

	if err4 != nil {
		t.Errorf("Some weird shit going on here: %s\n", err4)
	} else {
		fmt.Printf("Success!  Returned ApiKey: %s", apiKey)
	}
	//t.Errorf("hooray it returny.  First string: %s",str)
	//t.Errorf("Second string: %s", str2)

}
