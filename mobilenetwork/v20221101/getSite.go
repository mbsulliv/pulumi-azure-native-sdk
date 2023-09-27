// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets information about the specified mobile network site.
func LookupSite(ctx *pulumi.Context, args *LookupSiteArgs, opts ...pulumi.InvokeOption) (*LookupSiteResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSiteResult
	err := ctx.Invoke("azure-native:mobilenetwork/v20221101:getSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteArgs struct {
	// The name of the mobile network.
	MobileNetworkName string `pulumi:"mobileNetworkName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the mobile network site.
	SiteName string `pulumi:"siteName"`
}

// Site resource. Must be created in the same location as its parent mobile network.
type LookupSiteResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// An array of IDs of the network functions deployed in the site. Deleting the site will delete any network functions that are deployed in the site.
	NetworkFunctions []SubResourceResponse `pulumi:"networkFunctions"`
	// The provisioning state of the site resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupSiteOutput(ctx *pulumi.Context, args LookupSiteOutputArgs, opts ...pulumi.InvokeOption) LookupSiteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSiteResult, error) {
			args := v.(LookupSiteArgs)
			r, err := LookupSite(ctx, &args, opts...)
			var s LookupSiteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSiteResultOutput)
}

type LookupSiteOutputArgs struct {
	// The name of the mobile network.
	MobileNetworkName pulumi.StringInput `pulumi:"mobileNetworkName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the mobile network site.
	SiteName pulumi.StringInput `pulumi:"siteName"`
}

func (LookupSiteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteArgs)(nil)).Elem()
}

// Site resource. Must be created in the same location as its parent mobile network.
type LookupSiteResultOutput struct{ *pulumi.OutputState }

func (LookupSiteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteResult)(nil)).Elem()
}

func (o LookupSiteResultOutput) ToLookupSiteResultOutput() LookupSiteResultOutput {
	return o
}

func (o LookupSiteResultOutput) ToLookupSiteResultOutputWithContext(ctx context.Context) LookupSiteResultOutput {
	return o
}

func (o LookupSiteResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSiteResult] {
	return pulumix.Output[LookupSiteResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSiteResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupSiteResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSiteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.Name }).(pulumi.StringOutput)
}

// An array of IDs of the network functions deployed in the site. Deleting the site will delete any network functions that are deployed in the site.
func (o LookupSiteResultOutput) NetworkFunctions() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupSiteResult) []SubResourceResponse { return v.NetworkFunctions }).(SubResourceResponseArrayOutput)
}

// The provisioning state of the site resource.
func (o LookupSiteResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSiteResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSiteResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupSiteResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSiteResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSiteResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSiteResultOutput{})
}
