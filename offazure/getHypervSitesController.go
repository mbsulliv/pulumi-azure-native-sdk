// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package offazure

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a HypervSite
// Azure REST API version: 2023-06-06.
//
// Other available API versions: 2023-10-01-preview.
func LookupHypervSitesController(ctx *pulumi.Context, args *LookupHypervSitesControllerArgs, opts ...pulumi.InvokeOption) (*LookupHypervSitesControllerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupHypervSitesControllerResult
	err := ctx.Invoke("azure-native:offazure:getHypervSitesController", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHypervSitesControllerArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Site name
	SiteName string `pulumi:"siteName"`
}

// A HyperV SiteResource
type LookupHypervSitesControllerResult struct {
	// Gets or sets the on-premises agent details.
	AgentDetails *SiteAgentPropertiesResponse `pulumi:"agentDetails"`
	// Gets or sets the Appliance Name.
	ApplianceName *string `pulumi:"applianceName"`
	// Gets or sets the ARM ID of migration hub solution for SDS.
	DiscoverySolutionId *string `pulumi:"discoverySolutionId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Gets the Master Site this site is linked to.
	MasterSiteId string `pulumi:"masterSiteId"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Gets the service endpoint.
	ServiceEndpoint string `pulumi:"serviceEndpoint"`
	// Gets or sets the service principal identity details used by agent for
	// communication
	//             to the service.
	ServicePrincipalIdentityDetails *SiteSpnPropertiesResponse `pulumi:"servicePrincipalIdentityDetails"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupHypervSitesControllerOutput(ctx *pulumi.Context, args LookupHypervSitesControllerOutputArgs, opts ...pulumi.InvokeOption) LookupHypervSitesControllerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHypervSitesControllerResult, error) {
			args := v.(LookupHypervSitesControllerArgs)
			r, err := LookupHypervSitesController(ctx, &args, opts...)
			var s LookupHypervSitesControllerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHypervSitesControllerResultOutput)
}

type LookupHypervSitesControllerOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Site name
	SiteName pulumi.StringInput `pulumi:"siteName"`
}

func (LookupHypervSitesControllerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHypervSitesControllerArgs)(nil)).Elem()
}

// A HyperV SiteResource
type LookupHypervSitesControllerResultOutput struct{ *pulumi.OutputState }

func (LookupHypervSitesControllerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHypervSitesControllerResult)(nil)).Elem()
}

func (o LookupHypervSitesControllerResultOutput) ToLookupHypervSitesControllerResultOutput() LookupHypervSitesControllerResultOutput {
	return o
}

func (o LookupHypervSitesControllerResultOutput) ToLookupHypervSitesControllerResultOutputWithContext(ctx context.Context) LookupHypervSitesControllerResultOutput {
	return o
}

// Gets or sets the on-premises agent details.
func (o LookupHypervSitesControllerResultOutput) AgentDetails() SiteAgentPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupHypervSitesControllerResult) *SiteAgentPropertiesResponse { return v.AgentDetails }).(SiteAgentPropertiesResponsePtrOutput)
}

// Gets or sets the Appliance Name.
func (o LookupHypervSitesControllerResultOutput) ApplianceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHypervSitesControllerResult) *string { return v.ApplianceName }).(pulumi.StringPtrOutput)
}

// Gets or sets the ARM ID of migration hub solution for SDS.
func (o LookupHypervSitesControllerResultOutput) DiscoverySolutionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHypervSitesControllerResult) *string { return v.DiscoverySolutionId }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupHypervSitesControllerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHypervSitesControllerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupHypervSitesControllerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHypervSitesControllerResult) string { return v.Location }).(pulumi.StringOutput)
}

// Gets the Master Site this site is linked to.
func (o LookupHypervSitesControllerResultOutput) MasterSiteId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHypervSitesControllerResult) string { return v.MasterSiteId }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupHypervSitesControllerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHypervSitesControllerResult) string { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o LookupHypervSitesControllerResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHypervSitesControllerResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Gets the service endpoint.
func (o LookupHypervSitesControllerResultOutput) ServiceEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHypervSitesControllerResult) string { return v.ServiceEndpoint }).(pulumi.StringOutput)
}

// Gets or sets the service principal identity details used by agent for
// communication
//
//	to the service.
func (o LookupHypervSitesControllerResultOutput) ServicePrincipalIdentityDetails() SiteSpnPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupHypervSitesControllerResult) *SiteSpnPropertiesResponse {
		return v.ServicePrincipalIdentityDetails
	}).(SiteSpnPropertiesResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupHypervSitesControllerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupHypervSitesControllerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupHypervSitesControllerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupHypervSitesControllerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupHypervSitesControllerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHypervSitesControllerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHypervSitesControllerResultOutput{})
}
