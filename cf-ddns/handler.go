package function

import (
	"encoding/json"
	"log"
	"os"
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

func Handle(req []byte) string {
	var err error

	// FIXME: change to secrets
	apiKey := os.Getenv("API_KEY")
	email := os.Getenv("EMAIL")

	var request Request
	err = json.Unmarshal(req, &request)
	if err != nil {
		return "Unmarshall error"
	}

	ipService := NewIPService()
	dnsService, err := NewDNSService(apiKey, email)
	if err != nil {
		return "DNSService initialization error"
	}

	ddnsService := NewDDNSService(ipService, dnsService)

	err = ddnsService.UpdateDNSRecordIfNecessary(request.ZoneName, request.Domains)
	if err != nil {
		return "Update failed"
	}
	log.Println("Update DNS Record Finish!")

	return "Update complete"

}
