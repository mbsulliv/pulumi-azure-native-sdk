// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description for Gets the details of a static site.
func LookupStaticSite(ctx *pulumi.Context, args *LookupStaticSiteArgs, opts ...pulumi.InvokeOption) (*LookupStaticSiteResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupStaticSiteResult
	err := ctx.Invoke("azure-native:web/v20220901:getStaticSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStaticSiteArgs struct {
	// Name of the static site.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Static Site ARM resource.
type LookupStaticSiteResult struct {
	// <code>false</code> if config file is locked for this static web app; otherwise, <code>true</code>.
	AllowConfigFileUpdates *bool `pulumi:"allowConfigFileUpdates"`
	// The target branch in the repository.
	Branch *string `pulumi:"branch"`
	// Build properties to configure on the repository.
	BuildProperties *StaticSiteBuildPropertiesResponse `pulumi:"buildProperties"`
	// The content distribution endpoint for the static site.
	ContentDistributionEndpoint string `pulumi:"contentDistributionEndpoint"`
	// The custom domains associated with this static site.
	CustomDomains []string `pulumi:"customDomains"`
	// Database connections for the static site
	DatabaseConnections []DatabaseConnectionOverviewResponse `pulumi:"databaseConnections"`
	// The default autogenerated hostname for the static site.
	DefaultHostname string `pulumi:"defaultHostname"`
	// State indicating the status of the enterprise grade CDN serving traffic to the static web app.
	EnterpriseGradeCdnStatus *string `pulumi:"enterpriseGradeCdnStatus"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Managed service identity.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// Identity to use for Key Vault Reference authentication.
	KeyVaultReferenceIdentity string `pulumi:"keyVaultReferenceIdentity"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Backends linked to the static side
	LinkedBackends []StaticSiteLinkedBackendResponse `pulumi:"linkedBackends"`
	// Resource Location.
	Location string `pulumi:"location"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Private endpoint connections
	PrivateEndpointConnections []ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// The provider that submitted the last deployment to the primary environment of the static site.
	Provider *string `pulumi:"provider"`
	// State indicating whether public traffic are allowed or not for a static web app. Allowed Values: 'Enabled', 'Disabled' or an empty string.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// A user's github repository token. This is used to setup the Github Actions workflow file and API secrets.
	RepositoryToken *string `pulumi:"repositoryToken"`
	// URL for the repository of the static site.
	RepositoryUrl *string `pulumi:"repositoryUrl"`
	// Description of a SKU for a scalable resource.
	Sku *SkuDescriptionResponse `pulumi:"sku"`
	// State indicating whether staging environments are allowed or not allowed for a static web app.
	StagingEnvironmentPolicy *string `pulumi:"stagingEnvironmentPolicy"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Template options for generating a new repository.
	TemplateProperties *StaticSiteTemplateOptionsResponse `pulumi:"templateProperties"`
	// Resource type.
	Type string `pulumi:"type"`
	// User provided function apps registered with the static site
	UserProvidedFunctionApps []StaticSiteUserProvidedFunctionAppResponse `pulumi:"userProvidedFunctionApps"`
}

func LookupStaticSiteOutput(ctx *pulumi.Context, args LookupStaticSiteOutputArgs, opts ...pulumi.InvokeOption) LookupStaticSiteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStaticSiteResult, error) {
			args := v.(LookupStaticSiteArgs)
			r, err := LookupStaticSite(ctx, &args, opts...)
			var s LookupStaticSiteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStaticSiteResultOutput)
}

type LookupStaticSiteOutputArgs struct {
	// Name of the static site.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupStaticSiteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteArgs)(nil)).Elem()
}

// Static Site ARM resource.
type LookupStaticSiteResultOutput struct{ *pulumi.OutputState }

func (LookupStaticSiteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteResult)(nil)).Elem()
}

func (o LookupStaticSiteResultOutput) ToLookupStaticSiteResultOutput() LookupStaticSiteResultOutput {
	return o
}

func (o LookupStaticSiteResultOutput) ToLookupStaticSiteResultOutputWithContext(ctx context.Context) LookupStaticSiteResultOutput {
	return o
}

// <code>false</code> if config file is locked for this static web app; otherwise, <code>true</code>.
func (o LookupStaticSiteResultOutput) AllowConfigFileUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *bool { return v.AllowConfigFileUpdates }).(pulumi.BoolPtrOutput)
}

// The target branch in the repository.
func (o LookupStaticSiteResultOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

// Build properties to configure on the repository.
func (o LookupStaticSiteResultOutput) BuildProperties() StaticSiteBuildPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *StaticSiteBuildPropertiesResponse { return v.BuildProperties }).(StaticSiteBuildPropertiesResponsePtrOutput)
}

// The content distribution endpoint for the static site.
func (o LookupStaticSiteResultOutput) ContentDistributionEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.ContentDistributionEndpoint }).(pulumi.StringOutput)
}

// The custom domains associated with this static site.
func (o LookupStaticSiteResultOutput) CustomDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) []string { return v.CustomDomains }).(pulumi.StringArrayOutput)
}

// Database connections for the static site
func (o LookupStaticSiteResultOutput) DatabaseConnections() DatabaseConnectionOverviewResponseArrayOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) []DatabaseConnectionOverviewResponse { return v.DatabaseConnections }).(DatabaseConnectionOverviewResponseArrayOutput)
}

// The default autogenerated hostname for the static site.
func (o LookupStaticSiteResultOutput) DefaultHostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.DefaultHostname }).(pulumi.StringOutput)
}

// State indicating the status of the enterprise grade CDN serving traffic to the static web app.
func (o LookupStaticSiteResultOutput) EnterpriseGradeCdnStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *string { return v.EnterpriseGradeCdnStatus }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupStaticSiteResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.Id }).(pulumi.StringOutput)
}

// Managed service identity.
func (o LookupStaticSiteResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Identity to use for Key Vault Reference authentication.
func (o LookupStaticSiteResultOutput) KeyVaultReferenceIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.KeyVaultReferenceIdentity }).(pulumi.StringOutput)
}

// Kind of resource.
func (o LookupStaticSiteResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Backends linked to the static side
func (o LookupStaticSiteResultOutput) LinkedBackends() StaticSiteLinkedBackendResponseArrayOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) []StaticSiteLinkedBackendResponse { return v.LinkedBackends }).(StaticSiteLinkedBackendResponseArrayOutput)
}

// Resource Location.
func (o LookupStaticSiteResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource Name.
func (o LookupStaticSiteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.Name }).(pulumi.StringOutput)
}

// Private endpoint connections
func (o LookupStaticSiteResultOutput) PrivateEndpointConnections() ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) []ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput)
}

// The provider that submitted the last deployment to the primary environment of the static site.
func (o LookupStaticSiteResultOutput) Provider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *string { return v.Provider }).(pulumi.StringPtrOutput)
}

// State indicating whether public traffic are allowed or not for a static web app. Allowed Values: 'Enabled', 'Disabled' or an empty string.
func (o LookupStaticSiteResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// A user's github repository token. This is used to setup the Github Actions workflow file and API secrets.
func (o LookupStaticSiteResultOutput) RepositoryToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *string { return v.RepositoryToken }).(pulumi.StringPtrOutput)
}

// URL for the repository of the static site.
func (o LookupStaticSiteResultOutput) RepositoryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *string { return v.RepositoryUrl }).(pulumi.StringPtrOutput)
}

// Description of a SKU for a scalable resource.
func (o LookupStaticSiteResultOutput) Sku() SkuDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *SkuDescriptionResponse { return v.Sku }).(SkuDescriptionResponsePtrOutput)
}

// State indicating whether staging environments are allowed or not allowed for a static web app.
func (o LookupStaticSiteResultOutput) StagingEnvironmentPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *string { return v.StagingEnvironmentPolicy }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o LookupStaticSiteResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Template options for generating a new repository.
func (o LookupStaticSiteResultOutput) TemplateProperties() StaticSiteTemplateOptionsResponsePtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *StaticSiteTemplateOptionsResponse { return v.TemplateProperties }).(StaticSiteTemplateOptionsResponsePtrOutput)
}

// Resource type.
func (o LookupStaticSiteResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.Type }).(pulumi.StringOutput)
}

// User provided function apps registered with the static site
func (o LookupStaticSiteResultOutput) UserProvidedFunctionApps() StaticSiteUserProvidedFunctionAppResponseArrayOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) []StaticSiteUserProvidedFunctionAppResponse {
		return v.UserProvidedFunctionApps
	}).(StaticSiteUserProvidedFunctionAppResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStaticSiteResultOutput{})
}
