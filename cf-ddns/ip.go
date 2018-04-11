package function

import (
	"io/ioutil"
	"net/http"
)

type IPFetcher struct{}

func NewIPService() IPService {

	return &IPFetcher{}
}

func (f *IPFetcher) FetchIP() (string, error) {
	res, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "", err
	}

	ip, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}
