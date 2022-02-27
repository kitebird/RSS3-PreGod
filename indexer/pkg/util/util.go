package util

import (
	"time"

	"github.com/go-resty/resty/v2"
)

func Get(url string, headers map[string]string) ([]byte, error) {
	// Create a Resty Client
	client := resty.New()
	client.SetTimeout(1 * time.Second * 10)

	headers["User-Agent"] = "RSS3-PreGod"
	request := client.R().EnableTrace().SetHeaders(headers)

	// Get url
	resp, err := request.Get(url)

	return resp.Body(), err
}

func Post(url string, headers map[string]string, data string) (*resty.Response, error) {
	// Create a Resty Client
	client := resty.New()
	client.SetTimeout(1 * time.Second * 10)

	headers["User-Agent"] = "RSS3-PreGod"
	request := client.R().EnableTrace().SetHeaders(headers).SetBody(data)

	// Get url
	resp, err := request.Post(url)

	return resp, err
}
