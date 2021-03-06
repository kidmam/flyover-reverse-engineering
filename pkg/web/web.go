package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get fetches the contents of a URL
func Get(url string) (data []byte, err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer func() {
		err = res.Body.Close()
	}()
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("http status %d", res.StatusCode)
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	data = body
	return
}
