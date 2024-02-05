// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package offazure

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Method to get a site.
// Azure REST API version: 2023-06-06.
//
// Other available API versions: 2023-10-01-preview.
func LookupWebAppSitesController(ctx *pulumi.Context, args *LookupWebAppSitesControllerArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSitesControllerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppSitesControllerResult
	err := ctx.Invoke("azure-native:offazure:getWebAppSitesController", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppSitesControllerArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Site name
	SiteName string `pulumi:"siteName"`
	// Web app site name.
	WebAppSiteName string `pulumi:"webAppSiteName"`
}

// WebApp site web model.
type LookupWebAppSitesControllerResult struct {
	// Gets or sets the discovery scenario.
	DiscoveryScenario *string `pulumi:"discoveryScenario"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// provisioning state enum
	ProvisioningState string `pulumi:"provisioningState"`
	// Gets the service endpoint.
	ServiceEndpoint string `pulumi:"serviceEndpoint"`
	// Gets or sets the appliance details used by service to communicate
	//
	// to the appliance.
	SiteAppliancePropertiesCollection []SiteAppliancePropertiesResponse `pulumi:"siteAppliancePropertiesCollection"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupWebAppSitesControllerOutput(ctx *pulumi.Context, args LookupWebAppSitesControllerOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppSitesControllerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppSitesControllerResult, error) {
			args := v.(LookupWebAppSitesControllerArgs)
			r, err := LookupWebAppSitesController(ctx, &args, opts...)
			var s LookupWebAppSitesControllerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppSitesControllerResultOutput)
}

type LookupWebAppSitesControllerOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Site name
	SiteName pulumi.StringInput `pulumi:"siteName"`
	// Web app site name.
	WebAppSiteName pulumi.StringInput `pulumi:"webAppSiteName"`
}

func (LookupWebAppSitesControllerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSitesControllerArgs)(nil)).Elem()
}

// WebApp site web model.
type LookupWebAppSitesControllerResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppSitesControllerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSitesControllerResult)(nil)).Elem()
}

func (o LookupWebAppSitesControllerResultOutput) ToLookupWebAppSitesControllerResultOutput() LookupWebAppSitesControllerResultOutput {
	return o
}

func (o LookupWebAppSitesControllerResultOutput) ToLookupWebAppSitesControllerResultOutputWithContext(ctx context.Context) LookupWebAppSitesControllerResultOutput {
	return o
}

// Gets or sets the discovery scenario.
func (o LookupWebAppSitesControllerResultOutput) DiscoveryScenario() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSitesControllerResult) *string { return v.DiscoveryScenario }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWebAppSitesControllerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSitesControllerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWebAppSitesControllerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSitesControllerResult) string { return v.Name }).(pulumi.StringOutput)
}

// provisioning state enum
func (o LookupWebAppSitesControllerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSitesControllerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets the service endpoint.
func (o LookupWebAppSitesControllerResultOutput) ServiceEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSitesControllerResult) string { return v.ServiceEndpoint }).(pulumi.StringOutput)
}

// Gets or sets the appliance details used by service to communicate
//
// to the appliance.
func (o LookupWebAppSitesControllerResultOutput) SiteAppliancePropertiesCollection() SiteAppliancePropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupWebAppSitesControllerResult) []SiteAppliancePropertiesResponse {
		return v.SiteAppliancePropertiesCollection
	}).(SiteAppliancePropertiesResponseArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupWebAppSitesControllerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWebAppSitesControllerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWebAppSitesControllerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSitesControllerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppSitesControllerResultOutput{})
}
