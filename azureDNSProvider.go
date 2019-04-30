package azureDNS

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/dns/mgmt/2017-10-01/dns"
	"github.com/StackExchange/dnscontrol/models"
	"github.com/StackExchange/dnscontrol/providers"
	"github.com/pkg/errors"
)

/*
Azure DNS provider:
Info required in `creds.json`:
   - subscriptionId
   - clientId
   - clientSecret
   - tenantID

*/

type AzureDNSProvider struct {
	recordSetsClientAPI dns.RecordSetsClient
	zonesClientAPI      dns.ZonesClient
}

func newAzureDNSsp(conf map[string]string, metadata json.RawMessage) (providers.DNSServiceProvider, error) {
	return newAzureDns(conf, metadata)
}

func newAzureDns(m map[string]string, metadata json.RawMessage) (*AzureDNSProvider, error) {
	subscriptionId, clientId, clientSecret, tenantID := m["subscriptionId"], m["clientId"], m["clientSecret"], m["tenantID"]

	if subscriptionId == "" || clientId == "" || clientSecret == "" || tenantID == "" {
		return nil, errors.Errorf("You must provide the following configuration: \n SubscriptionID, ClientID, ClientSecret, TenantID")
	}

	api := &AzureDNSProvider{recordSetsClientAPI: dns.NewRecordSetsClient(subscriptionId), zonesClientAPI: dns.NewZonesClient(subscriptionId)}

	err := api.getZones()
	if err != nil {
		return nil, err
	}
	return api, nil
}

func (c *AzureDNSProvider) GetDomainCorrections(dc *models.DomainConfig) ([]*models.Correction, error) {

	return nil, nil
}

var features = providers.DocumentationNotes{
	providers.CanUseAlias:            providers.Can(),
	providers.DocCreateDomains:       providers.Can(),
	providers.DocOfficiallySupported: providers.Cannot(""),
	providers.CanUsePTR:              providers.Can(),
	providers.CanUseSRV:              providers.Can(),
	providers.CanUseTXTMulti:         providers.Can(),
	providers.CanUseCAA:              providers.Can(),
}

func (c *AzureDNSProvider) GetNameservers(domain string) ([]*models.Nameserver, error) {
	// if c.domainIndex == nil {
	// 	if err := c.fetchDomainList(); err != nil {
	// 		return nil, err
	// 	}
	// }
	// ns, ok := c.nameservers[domain]
	// if !ok {
	// 	return nil, errors.Errorf("Nameservers for %s not found in cloudflare account", domain)
	// }
	// return models.StringsToNameservers(ns), nil
	return nil, nil
}

func (c *AzureDNSProvider) getZones() error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)

	defer cancel()

	top := int32(10)

	_, error := c.zonesClientAPI.List(ctx, &top)

	return error
}

func init() {
	providers.RegisterDomainServiceProviderType("AzureDNS", newAzureDNSsp, features)
}

func main() {

}
