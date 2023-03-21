package httprequest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func RequestGET(endpoint string, path string, port int) {
	reqURL := fmt.Sprintf("https://%s:%d%s", endpoint, port, path)
	res, err := http.Get(reqURL)
	if err != nil {
		fmt.Printf("error http req: %s\n", err)
		os.Exit(1)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Go cha!! response code: %d\n", res.StatusCode)
	fmt.Printf("client: response body: %s\n", resBody)
}
