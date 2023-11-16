// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230606

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Method to get a site.
func LookupSqlSitesController(ctx *pulumi.Context, args *LookupSqlSitesControllerArgs, opts ...pulumi.InvokeOption) (*LookupSqlSitesControllerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSqlSitesControllerResult
	err := ctx.Invoke("azure-native:offazure/v20230606:getSqlSitesController", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlSitesControllerArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Site name
	SiteName string `pulumi:"siteName"`
	// SQL site name.
	SqlSiteName string `pulumi:"sqlSiteName"`
}

// SQL site web model.
type LookupSqlSitesControllerResult struct {
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

func LookupSqlSitesControllerOutput(ctx *pulumi.Context, args LookupSqlSitesControllerOutputArgs, opts ...pulumi.InvokeOption) LookupSqlSitesControllerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlSitesControllerResult, error) {
			args := v.(LookupSqlSitesControllerArgs)
			r, err := LookupSqlSitesController(ctx, &args, opts...)
			var s LookupSqlSitesControllerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlSitesControllerResultOutput)
}

type LookupSqlSitesControllerOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Site name
	SiteName pulumi.StringInput `pulumi:"siteName"`
	// SQL site name.
	SqlSiteName pulumi.StringInput `pulumi:"sqlSiteName"`
}

func (LookupSqlSitesControllerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlSitesControllerArgs)(nil)).Elem()
}

// SQL site web model.
type LookupSqlSitesControllerResultOutput struct{ *pulumi.OutputState }

func (LookupSqlSitesControllerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlSitesControllerResult)(nil)).Elem()
}

func (o LookupSqlSitesControllerResultOutput) ToLookupSqlSitesControllerResultOutput() LookupSqlSitesControllerResultOutput {
	return o
}

func (o LookupSqlSitesControllerResultOutput) ToLookupSqlSitesControllerResultOutputWithContext(ctx context.Context) LookupSqlSitesControllerResultOutput {
	return o
}

// Gets or sets the discovery scenario.
func (o LookupSqlSitesControllerResultOutput) DiscoveryScenario() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlSitesControllerResult) *string { return v.DiscoveryScenario }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSqlSitesControllerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlSitesControllerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSqlSitesControllerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlSitesControllerResult) string { return v.Name }).(pulumi.StringOutput)
}

// provisioning state enum
func (o LookupSqlSitesControllerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlSitesControllerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets the service endpoint.
func (o LookupSqlSitesControllerResultOutput) ServiceEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlSitesControllerResult) string { return v.ServiceEndpoint }).(pulumi.StringOutput)
}

// Gets or sets the appliance details used by service to communicate
//
// to the appliance.
func (o LookupSqlSitesControllerResultOutput) SiteAppliancePropertiesCollection() SiteAppliancePropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupSqlSitesControllerResult) []SiteAppliancePropertiesResponse {
		return v.SiteAppliancePropertiesCollection
	}).(SiteAppliancePropertiesResponseArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSqlSitesControllerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSqlSitesControllerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSqlSitesControllerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlSitesControllerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlSitesControllerResultOutput{})
}
