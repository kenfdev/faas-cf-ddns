package function

type DDNSService struct {
	ipService  IPService
	dnsService DNSService
}

func NewDDNSService(ipService IPService, dnsService DNSService) *DDNSService {
	return &DDNSService{
		ipService:  ipService,
		dnsService: dnsService,
	}
}

func (d *DDNSService) UpdateDNSRecordIfNecessary(zoneName string, domains []string) error {

	ip, err := d.ipService.FetchIP()
	if err != nil {
		// Fetch IP Error
		return err
	}

	rec, err := d.dnsService.FetchDNSRecordFor(zoneName, domains[0])
	if err != nil {
		// "Fetching DNS record error"
		return err
	}

	if ip == rec.Content {
		// "Nothing to change"
		return nil
	}

	err = d.dnsService.UpdateDNSRecordIPFor(rec.ZoneID, domains, ip)
	if err != nil {
		// "DNS record update error"
		return err
	}
	return nil
}
