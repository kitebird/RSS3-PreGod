package util

import "github.com/go-resty/resty/v2"

func GetURL(url string, headers map[string]string) ([]byte, error) {
	// Create a Resty Client
	client := resty.New()
	request := client.R().EnableTrace().SetHeaders(headers)

	// GetURL url
	resp, err := request.Get(url)

	return resp.Body(), err
}
