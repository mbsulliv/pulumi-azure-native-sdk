// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package offazure

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a MasterSite
// Azure REST API version: 2023-06-06.
//
// Other available API versions: 2023-10-01-preview.
func LookupMasterSitesController(ctx *pulumi.Context, args *LookupMasterSitesControllerArgs, opts ...pulumi.InvokeOption) (*LookupMasterSitesControllerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMasterSitesControllerResult
	err := ctx.Invoke("azure-native:offazure:getMasterSitesController", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMasterSitesControllerArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Site name
	SiteName string `pulumi:"siteName"`
}

// A MasterSite
type LookupMasterSitesControllerResult struct {
	// Gets or sets a value indicating whether multiple sites per site type are
	// allowed.
	AllowMultipleSites *bool `pulumi:"allowMultipleSites"`
	// Gets or sets a value for customer storage account ARM id.
	CustomerStorageAccountArmId *string `pulumi:"customerStorageAccountArmId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Gets the nested sites under Master Site.
	NestedSites []string `pulumi:"nestedSites"`
	// Gets the private endpoint connections.
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// provisioning state enum
	ProvisioningState string `pulumi:"provisioningState"`
	// Gets or sets the state of public network access.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Gets or sets the sites that are a part of Master Site.
	//             The key
	// should contain the Site ARM name.
	Sites []string `pulumi:"sites"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupMasterSitesControllerOutput(ctx *pulumi.Context, args LookupMasterSitesControllerOutputArgs, opts ...pulumi.InvokeOption) LookupMasterSitesControllerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMasterSitesControllerResult, error) {
			args := v.(LookupMasterSitesControllerArgs)
			r, err := LookupMasterSitesController(ctx, &args, opts...)
			var s LookupMasterSitesControllerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMasterSitesControllerResultOutput)
}

type LookupMasterSitesControllerOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Site name
	SiteName pulumi.StringInput `pulumi:"siteName"`
}

func (LookupMasterSitesControllerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMasterSitesControllerArgs)(nil)).Elem()
}

// A MasterSite
type LookupMasterSitesControllerResultOutput struct{ *pulumi.OutputState }

func (LookupMasterSitesControllerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMasterSitesControllerResult)(nil)).Elem()
}

func (o LookupMasterSitesControllerResultOutput) ToLookupMasterSitesControllerResultOutput() LookupMasterSitesControllerResultOutput {
	return o
}

func (o LookupMasterSitesControllerResultOutput) ToLookupMasterSitesControllerResultOutputWithContext(ctx context.Context) LookupMasterSitesControllerResultOutput {
	return o
}

// Gets or sets a value indicating whether multiple sites per site type are
// allowed.
func (o LookupMasterSitesControllerResultOutput) AllowMultipleSites() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupMasterSitesControllerResult) *bool { return v.AllowMultipleSites }).(pulumi.BoolPtrOutput)
}

// Gets or sets a value for customer storage account ARM id.
func (o LookupMasterSitesControllerResultOutput) CustomerStorageAccountArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMasterSitesControllerResult) *string { return v.CustomerStorageAccountArmId }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupMasterSitesControllerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMasterSitesControllerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupMasterSitesControllerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMasterSitesControllerResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupMasterSitesControllerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMasterSitesControllerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets the nested sites under Master Site.
func (o LookupMasterSitesControllerResultOutput) NestedSites() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMasterSitesControllerResult) []string { return v.NestedSites }).(pulumi.StringArrayOutput)
}

// Gets the private endpoint connections.
func (o LookupMasterSitesControllerResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupMasterSitesControllerResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// provisioning state enum
func (o LookupMasterSitesControllerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMasterSitesControllerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets or sets the state of public network access.
func (o LookupMasterSitesControllerResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMasterSitesControllerResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Gets or sets the sites that are a part of Master Site.
//
//	The key
//
// should contain the Site ARM name.
func (o LookupMasterSitesControllerResultOutput) Sites() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMasterSitesControllerResult) []string { return v.Sites }).(pulumi.StringArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupMasterSitesControllerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMasterSitesControllerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupMasterSitesControllerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMasterSitesControllerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupMasterSitesControllerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMasterSitesControllerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMasterSitesControllerResultOutput{})
}
