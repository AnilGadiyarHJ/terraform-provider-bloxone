package provider

import (
	"context"
	"fmt"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/service/anycast"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	bloxoneclient "github.com/infobloxopen/bloxone-go-client/client"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/service/dns_config"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/service/dns_data"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/service/fw"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/service/infra_mgmt"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/service/infra_provision"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/service/ipam"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/service/keys"
)

// Ensure BloxOneProvider satisfies various provider interfaces.
var _ provider.Provider = &BloxOneProvider{}

// BloxOneProvider defines the provider implementation.
type BloxOneProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
	commit  string
}

// BloxOneProviderModel describes the provider data model.
type BloxOneProviderModel struct {
	CSPUrl types.String `tfsdk:"csp_url"`
	APIKey types.String `tfsdk:"api_key"`
}

func (p *BloxOneProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "bloxone"
	resp.Version = p.version
}

func (p *BloxOneProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The BloxOne provider is used to interact with the resources supported by Infoblox BloxOne API.",
		Attributes: map[string]schema.Attribute{
			"csp_url": schema.StringAttribute{
				MarkdownDescription: "URL for BloxOne Cloud Services Portal. Can also be configured using the `BLOXONE_CSP_URL` environment variable.",
				Optional:            true,
			},
			"api_key": schema.StringAttribute{
				MarkdownDescription: "API key for accessing the BloxOne API. Can also be configured by using the `BLOXONE_API_KEY` environment variable. https://docs.infoblox.com/space/BloxOneCloud/35430405/Configuring+User+API+Keys",
				Optional:            true,
			},
		},
	}
}

func (p *BloxOneProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data BloxOneProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	client, err := bloxoneclient.NewAPIClient(bloxoneclient.Configuration{
		ClientName: fmt.Sprintf("terraform/%s#%s", p.version, p.commit),
		APIKey:     data.APIKey.ValueString(),
		CSPURL:     data.CSPUrl.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client error", fmt.Sprintf("Unable to create new API client: %s", err))
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *BloxOneProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		ipam.NewIpamHostResource,
		ipam.NewIpSpaceResource,
		ipam.NewSubnetResource,
		ipam.NewAddressBlockResource,
		ipam.NewAddressResource,
		ipam.NewRangeResource,
		ipam.NewFixedAddressResource,
		ipam.NewServerResource,
		ipam.NewHaGroupResource,
		ipam.NewOptionSpaceResource,
		ipam.NewOptionGroupResource,
		ipam.NewOptionCodeResource,

		dns_config.NewViewResource,
		dns_config.NewAuthNsgResource,
		dns_config.NewForwardZoneResource,
		dns_config.NewAuthZoneResource,
		dns_config.NewForwardNsgResource,
		dns_config.NewDelegationResource,
		dns_config.NewServerResource,
		dns_config.NewAclResource,

		dns_data.NewRecordAResource,
		dns_data.NewRecordAAAAResource,
		dns_data.NewRecordCNAMEResource,
		dns_data.NewRecordDNAMEResource,
		dns_data.NewRecordMXResource,
		dns_data.NewRecordNSResource,
		dns_data.NewRecordPTRResource,
		dns_data.NewRecordSRVResource,
		dns_data.NewRecordTXTResource,
		dns_data.NewRecordNAPTRResource,
		dns_data.NewRecordHTTPSResource,
		dns_data.NewRecordSVCBResource,
		dns_data.NewRecordCAAResource,
		dns_data.NewRecordGenericResource,

		infra_provision.NewUIJoinTokenResource,

		infra_mgmt.NewHostsResource,
		infra_mgmt.NewServicesResource,

		keys.NewTsigResource,

		anycast.NewOnPremAnycastManagerResource,
    
		fw.NewAccessCodesResource,
		fw.NewNamedListsResource,
		fw.NewNetworkListsResource,
		fw.NewInternalDomainListsResource,
	}
}

func (p *BloxOneProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		ipam.NewDhcpHostDataSource,
		ipam.NewIpamHostDataSource,
		ipam.NewIpSpaceDataSource,
		ipam.NewSubnetDataSource,
		ipam.NewAddressBlockDataSource,
		ipam.NewAddressDataSource,
		ipam.NewRangeDataSource,
		ipam.NewFixedAddressDataSource,
		ipam.NewServerDataSource,
		ipam.NewHaGroupDataSource,
		ipam.NewOptionCodeDataSource,
		ipam.NewOptionSpaceDataSource,
		ipam.NewOptionGroupDataSource,
		ipam.NewIpamNextAvailableIPDataSource,
		ipam.NewNextAvailableSubnetDataSource,
		ipam.NewNextAvailableAddressBlockDataSource,

		dns_config.NewViewDataSource,
		dns_config.NewAuthNsgDataSource,
		dns_config.NewHostDataSource,
		dns_config.NewForwardZoneDataSource,
		dns_config.NewAuthZoneDataSource,
		dns_config.NewForwardNsgDataSource,
		dns_config.NewDelegationDataSource,
		dns_config.NewServerDataSource,
		dns_config.NewAclDataSource,

		dns_data.NewRecordADataSource,
		dns_data.NewRecordAAAADataSource,
		dns_data.NewRecordCNAMEDataSource,
		dns_data.NewRecordDNAMEDataSource,
		dns_data.NewRecordMXDataSource,
		dns_data.NewRecordNSDataSource,
		dns_data.NewRecordPTRDataSource,
		dns_data.NewRecordSRVDataSource,
		dns_data.NewRecordTXTDataSource,
		dns_data.NewRecordNAPTRDataSource,
		dns_data.NewRecordHTTPSDataSource,
		dns_data.NewRecordSVCBDataSource,
		dns_data.NewRecordCAADataSource,
		dns_data.NewRecordGenericDataSource,

		infra_provision.NewUIJoinTokenDataSource,

		infra_mgmt.NewHostsDataSource,
		infra_mgmt.NewServicesDataSource,

		keys.NewTsigDataSource,
		keys.NewKerberosDataSource,

		anycast.NewOnPremAnycastManagerDataSource,

		fw.NewAccessCodesDataSource,
		fw.NewNamedListsDataSource,
		fw.NewNetworkListsDataSource,
		fw.NewInternalDomainListsDataSource,
		fw.NewPoPRegionsDataSource,
	}
}

func New(version, commit string) func() provider.Provider {
	return func() provider.Provider {
		return &BloxOneProvider{
			version: version,
			commit:  commit,
		}
	}
}
