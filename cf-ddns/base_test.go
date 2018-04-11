package function

var mockDNSService *MockDNSService
var mockIPService *MockIPService

func createDNSRecord() *DNSRecord {
	return &DNSRecord{
		ID:       "id",
		Type:     "A",
		Name:     "SomeName",
		Content:  "1.1.1.1",
		ZoneID:   "SomeZoneID",
		ZoneName: "SomeZoneName",
	}
}
