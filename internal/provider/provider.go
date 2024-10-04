package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = &NotionProvider{}

type NotionProvider struct {
}

type NotionProviderModel struct {
	NotionIntegrationToken string `tfsdk:"notion_integration_token"`
	NotionApiVersion       string `tfsdk:"notion_api_version"`
}

func New() provider.Provider {
	return &NotionProvider{}
}

func (p *NotionProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "notion"
}

func (p *NotionProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"notion_integration_token": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
			},
			"notion_api_version": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (p *NotionProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config NotionProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (p *NotionProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		// Return your data sources here
	}
}

func (p *NotionProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewNotionResource,
	}
}
