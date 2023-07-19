package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	//res, err := http.Get("https://blog.logrocket.com/making-http-requests-in-go/")
	req, err := http.NewRequest(http.MethodGet, "https://blog.logrocket.com/making-http-requests-in-go/", nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", resBody)

}
