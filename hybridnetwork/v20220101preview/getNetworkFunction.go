// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about the specified network function resource.
func LookupNetworkFunction(ctx *pulumi.Context, args *LookupNetworkFunctionArgs, opts ...pulumi.InvokeOption) (*LookupNetworkFunctionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkFunctionResult
	err := ctx.Invoke("azure-native:hybridnetwork/v20220101preview:getNetworkFunction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkFunctionArgs struct {
	// The name of the network function resource.
	NetworkFunctionName string `pulumi:"networkFunctionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Network function resource response.
type LookupNetworkFunctionResult struct {
	// The reference to the device resource. Once set, it cannot be updated.
	Device *SubResourceResponse `pulumi:"device"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The resource URI of the managed application.
	ManagedApplication SubResourceResponse `pulumi:"managedApplication"`
	// The parameters for the managed application.
	ManagedApplicationParameters interface{} `pulumi:"managedApplicationParameters"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The network function container configurations from the user.
	NetworkFunctionContainerConfigurations interface{} `pulumi:"networkFunctionContainerConfigurations"`
	// The network function configurations from the user.
	NetworkFunctionUserConfigurations []NetworkFunctionUserConfigurationResponse `pulumi:"networkFunctionUserConfigurations"`
	// The provisioning state of the network function resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The service key for the network function resource.
	ServiceKey string `pulumi:"serviceKey"`
	// The sku name for the network function. Once set, it cannot be updated.
	SkuName *string `pulumi:"skuName"`
	// The sku type for the network function.
	SkuType string `pulumi:"skuType"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The vendor name for the network function. Once set, it cannot be updated.
	VendorName *string `pulumi:"vendorName"`
	// The vendor provisioning state for the network function resource.
	VendorProvisioningState string `pulumi:"vendorProvisioningState"`
}

func LookupNetworkFunctionOutput(ctx *pulumi.Context, args LookupNetworkFunctionOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkFunctionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkFunctionResult, error) {
			args := v.(LookupNetworkFunctionArgs)
			r, err := LookupNetworkFunction(ctx, &args, opts...)
			var s LookupNetworkFunctionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkFunctionResultOutput)
}

type LookupNetworkFunctionOutputArgs struct {
	// The name of the network function resource.
	NetworkFunctionName pulumi.StringInput `pulumi:"networkFunctionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkFunctionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkFunctionArgs)(nil)).Elem()
}

// Network function resource response.
type LookupNetworkFunctionResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkFunctionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkFunctionResult)(nil)).Elem()
}

func (o LookupNetworkFunctionResultOutput) ToLookupNetworkFunctionResultOutput() LookupNetworkFunctionResultOutput {
	return o
}

func (o LookupNetworkFunctionResultOutput) ToLookupNetworkFunctionResultOutputWithContext(ctx context.Context) LookupNetworkFunctionResultOutput {
	return o
}

// The reference to the device resource. Once set, it cannot be updated.
func (o LookupNetworkFunctionResultOutput) Device() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) *SubResourceResponse { return v.Device }).(SubResourceResponsePtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupNetworkFunctionResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupNetworkFunctionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupNetworkFunctionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) string { return v.Location }).(pulumi.StringOutput)
}

// The resource URI of the managed application.
func (o LookupNetworkFunctionResultOutput) ManagedApplication() SubResourceResponseOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) SubResourceResponse { return v.ManagedApplication }).(SubResourceResponseOutput)
}

// The parameters for the managed application.
func (o LookupNetworkFunctionResultOutput) ManagedApplicationParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) interface{} { return v.ManagedApplicationParameters }).(pulumi.AnyOutput)
}

// The name of the resource
func (o LookupNetworkFunctionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The network function container configurations from the user.
func (o LookupNetworkFunctionResultOutput) NetworkFunctionContainerConfigurations() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) interface{} { return v.NetworkFunctionContainerConfigurations }).(pulumi.AnyOutput)
}

// The network function configurations from the user.
func (o LookupNetworkFunctionResultOutput) NetworkFunctionUserConfigurations() NetworkFunctionUserConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) []NetworkFunctionUserConfigurationResponse {
		return v.NetworkFunctionUserConfigurations
	}).(NetworkFunctionUserConfigurationResponseArrayOutput)
}

// The provisioning state of the network function resource.
func (o LookupNetworkFunctionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The service key for the network function resource.
func (o LookupNetworkFunctionResultOutput) ServiceKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) string { return v.ServiceKey }).(pulumi.StringOutput)
}

// The sku name for the network function. Once set, it cannot be updated.
func (o LookupNetworkFunctionResultOutput) SkuName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) *string { return v.SkuName }).(pulumi.StringPtrOutput)
}

// The sku type for the network function.
func (o LookupNetworkFunctionResultOutput) SkuType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) string { return v.SkuType }).(pulumi.StringOutput)
}

// The system meta data relating to this resource.
func (o LookupNetworkFunctionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupNetworkFunctionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupNetworkFunctionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) string { return v.Type }).(pulumi.StringOutput)
}

// The vendor name for the network function. Once set, it cannot be updated.
func (o LookupNetworkFunctionResultOutput) VendorName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) *string { return v.VendorName }).(pulumi.StringPtrOutput)
}

// The vendor provisioning state for the network function resource.
func (o LookupNetworkFunctionResultOutput) VendorProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) string { return v.VendorProvisioningState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkFunctionResultOutput{})
}
