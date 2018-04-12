package function

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/openfaas-incubator/go-function-sdk"
)

type DNSRecord struct {
	ID       string `json:"id,omitempty"`
	Type     string `json:"type,omitempty"`
	Name     string `json:"name,omitempty"`
	Content  string `json:"content,omitempty"`
	ZoneID   string `json:"zone_id,omitempty"`
	ZoneName string `json:"zone_name,omitempty"`
}

type DNSService interface {
	FetchDNSRecordFor(zoneName string, domainName string) (*DNSRecord, error)
	UpdateDNSRecordIPFor(zoneID string, domains []string, ip string) error
}

type IPService interface {
	FetchIP() (string, error)
}

type Request struct {
	ZoneName string   `json:"zone_name"`
	Domains  []string `json:"domains"`
}

type Response struct {
	Result string `json:"result"`
}

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error

	// FIXME: change to secrets
	apiKey := os.Getenv("API_KEY")
	email := os.Getenv("EMAIL")

	var request Request
	err = json.Unmarshal(req.Body, &request)
	if err != nil {
		return handler.Response{
			Body:       []byte("Unmarshall error"),
			StatusCode: http.StatusBadRequest,
		}, err
	}

	ipService := NewIPService()
	dnsService, err := NewDNSService(apiKey, email)
	if err != nil {
		return handler.Response{
			Body:       []byte("DNSService initialization error"),
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	ddnsService := NewDDNSService(ipService, dnsService)

	err = ddnsService.UpdateDNSRecordIfNecessary(request.ZoneName, request.Domains)
	if err != nil {
		return handler.Response{
			Body:       []byte("Update failed"),
			StatusCode: http.StatusInternalServerError,
		}, err
	}
	log.Println("Update DNS Record Finish!")

	return handler.Response{
		Body:       []byte("Update complete"),
		StatusCode: http.StatusOK,
	}, err

}
