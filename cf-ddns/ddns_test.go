package function

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupDDNS() *DDNSService {
	mockDNSService = &MockDNSService{}
	mockIPService = &MockIPService{}

	return &DDNSService{dnsService: mockDNSService, ipService: mockIPService}
}

func TestUpdateDNSRecordIfNecessaryDoesNotUpdateDNSIfIPUnChanged(t *testing.T) {
	// Arrange
	ddns := setupDDNS()
	zoneName := "example.com"
	domains := []string{"foo.example.com", "bar.example.com"}

	ip := "1.1.1.1"
	mockIPService.On("FetchIP").Return(ip, nil)

	dnsRec := createDNSRecord()
	dnsRec.Content = ip
	mockDNSService.On("FetchDNSRecordFor", zoneName, domains[0]).Return(dnsRec, nil)

	// Act
	err := ddns.UpdateDNSRecordIfNecessary(zoneName, domains)

	// Assert
	assert.Nil(t, err)
	mockDNSService.AssertNotCalled(t, "UpdateDNSRecordIPFor")

}

func TestUpdateDNSRecordIfNecessaryCallsUpdateDNSIfIPChanged(t *testing.T) {
	// Arrange
	ddns := setupDDNS()
	zoneName := "example.com"
	domains := []string{"foo.example.com", "bar.example.com"}

	ip := "1.1.1.1"
	mockIPService.On("FetchIP").Return(ip, nil)

	zoneID := "SomeZoneID"
	previousIP := "1.2.3.4"
	dnsRec := createDNSRecord()
	dnsRec.Content = previousIP
	dnsRec.ZoneID = zoneID
	mockDNSService.On("FetchDNSRecordFor", zoneName, domains[0]).Return(dnsRec, nil)
	mockDNSService.On("UpdateDNSRecordIPFor", zoneID, domains, ip).Return(nil)

	// Act
	err := ddns.UpdateDNSRecordIfNecessary(zoneName, domains)

	// Assert
	assert.Nil(t, err)
	mockDNSService.AssertExpectations(t)

}
