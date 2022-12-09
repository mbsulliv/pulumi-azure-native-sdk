// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Network slice resource. Must be created in the same location as its parent mobile network.
func LookupSlice(ctx *pulumi.Context, args *LookupSliceArgs, opts ...pulumi.InvokeOption) (*LookupSliceResult, error) {
	var rv LookupSliceResult
	err := ctx.Invoke("azure-native:mobilenetwork/v20221101:getSlice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSliceArgs struct {
	// The name of the mobile network.
	MobileNetworkName string `pulumi:"mobileNetworkName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the network slice.
	SliceName string `pulumi:"sliceName"`
}

// Network slice resource. Must be created in the same location as its parent mobile network.
type LookupSliceResult struct {
	// An optional description for this network slice.
	Description *string `pulumi:"description"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the network slice resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Single-network slice selection assistance information (S-NSSAI). Unique at the scope of a mobile network.
	Snssai SnssaiResponse `pulumi:"snssai"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupSliceOutput(ctx *pulumi.Context, args LookupSliceOutputArgs, opts ...pulumi.InvokeOption) LookupSliceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSliceResult, error) {
			args := v.(LookupSliceArgs)
			r, err := LookupSlice(ctx, &args, opts...)
			var s LookupSliceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSliceResultOutput)
}

type LookupSliceOutputArgs struct {
	// The name of the mobile network.
	MobileNetworkName pulumi.StringInput `pulumi:"mobileNetworkName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the network slice.
	SliceName pulumi.StringInput `pulumi:"sliceName"`
}

func (LookupSliceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSliceArgs)(nil)).Elem()
}

// Network slice resource. Must be created in the same location as its parent mobile network.
type LookupSliceResultOutput struct{ *pulumi.OutputState }

func (LookupSliceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSliceResult)(nil)).Elem()
}

func (o LookupSliceResultOutput) ToLookupSliceResultOutput() LookupSliceResultOutput {
	return o
}

func (o LookupSliceResultOutput) ToLookupSliceResultOutputWithContext(ctx context.Context) LookupSliceResultOutput {
	return o
}

// An optional description for this network slice.
func (o LookupSliceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSliceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSliceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSliceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupSliceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSliceResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSliceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSliceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the network slice resource.
func (o LookupSliceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSliceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Single-network slice selection assistance information (S-NSSAI). Unique at the scope of a mobile network.
func (o LookupSliceResultOutput) Snssai() SnssaiResponseOutput {
	return o.ApplyT(func(v LookupSliceResult) SnssaiResponse { return v.Snssai }).(SnssaiResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSliceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSliceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupSliceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSliceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSliceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSliceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSliceResultOutput{})
}
