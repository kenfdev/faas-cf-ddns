package function

import (
	cloudflare "github.com/cloudflare/cloudflare-go"
)

type CFWrapper struct {
	api *cloudflare.API
}

func NewDNSService(key string, email string) (DNSService, error) {
	api, err := cloudflare.New(key, email)
	if err != nil {
		return nil, err
	}

	return &CFWrapper{api: api}, nil
}

func (d *CFWrapper) FetchDNSRecordFor(zoneName string, domainName string) (*DNSRecord, error) {

	zoneID, err := d.api.ZoneIDByName(zoneName)
	if err != nil {
		return nil, err
	}

	r := cloudflare.DNSRecord{Name: domainName, Type: "A"}
	recs, err := d.api.DNSRecords(zoneID, r)
	if err != nil {
		return nil, err
	}

	rec := recs[0]
	result := &DNSRecord{
		ID:       rec.ID,
		Type:     rec.Type,
		Name:     rec.Name,
		Content:  rec.Content,
		ZoneID:   rec.ZoneID,
		ZoneName: rec.ZoneName,
	}

	return result, nil

}

func (d *CFWrapper) UpdateDNSRecordIPFor(zoneID string, domains []string, ip string) error {
	for _, domain := range domains {

		r := cloudflare.DNSRecord{Name: domain, Type: "A"}
		recs, err := d.api.DNSRecords(zoneID, r)
		if err != nil {
			return err
		}

		for _, rec := range recs {
			newRec := cloudflare.DNSRecord{
				Type:    rec.Type,
				Name:    rec.Name,
				Content: ip,
			}
			err := d.api.UpdateDNSRecord(zoneID, rec.ID, newRec)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
