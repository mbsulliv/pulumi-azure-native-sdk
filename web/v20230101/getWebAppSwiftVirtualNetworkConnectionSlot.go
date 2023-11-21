// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description for Gets a Swift Virtual Network connection.
func LookupWebAppSwiftVirtualNetworkConnectionSlot(ctx *pulumi.Context, args *LookupWebAppSwiftVirtualNetworkConnectionSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSwiftVirtualNetworkConnectionSlotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppSwiftVirtualNetworkConnectionSlotResult
	err := ctx.Invoke("azure-native:web/v20230101:getWebAppSwiftVirtualNetworkConnectionSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppSwiftVirtualNetworkConnectionSlotArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will get a gateway for the production slot's Virtual Network.
	Slot string `pulumi:"slot"`
}

// Swift Virtual Network Contract. This is used to enable the new Swift way of doing virtual network integration.
type LookupWebAppSwiftVirtualNetworkConnectionSlotResult struct {
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// The Virtual Network subnet's resource ID. This is the subnet that this Web App will join. This subnet must have a delegation to Microsoft.Web/serverFarms defined first.
	SubnetResourceId *string `pulumi:"subnetResourceId"`
	// A flag that specifies if the scale unit this Web App is on supports Swift integration.
	SwiftSupported *bool `pulumi:"swiftSupported"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupWebAppSwiftVirtualNetworkConnectionSlotOutput(ctx *pulumi.Context, args LookupWebAppSwiftVirtualNetworkConnectionSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppSwiftVirtualNetworkConnectionSlotResult, error) {
			args := v.(LookupWebAppSwiftVirtualNetworkConnectionSlotArgs)
			r, err := LookupWebAppSwiftVirtualNetworkConnectionSlot(ctx, &args, opts...)
			var s LookupWebAppSwiftVirtualNetworkConnectionSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput)
}

type LookupWebAppSwiftVirtualNetworkConnectionSlotOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will get a gateway for the production slot's Virtual Network.
	Slot pulumi.StringInput `pulumi:"slot"`
}

func (LookupWebAppSwiftVirtualNetworkConnectionSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSwiftVirtualNetworkConnectionSlotArgs)(nil)).Elem()
}

// Swift Virtual Network Contract. This is used to enable the new Swift way of doing virtual network integration.
type LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSwiftVirtualNetworkConnectionSlotResult)(nil)).Elem()
}

func (o LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput) ToLookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput() LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput {
	return o
}

func (o LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput) ToLookupWebAppSwiftVirtualNetworkConnectionSlotResultOutputWithContext(ctx context.Context) LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput {
	return o
}

// Resource Id.
func (o LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSwiftVirtualNetworkConnectionSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSwiftVirtualNetworkConnectionSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSwiftVirtualNetworkConnectionSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Virtual Network subnet's resource ID. This is the subnet that this Web App will join. This subnet must have a delegation to Microsoft.Web/serverFarms defined first.
func (o LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput) SubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSwiftVirtualNetworkConnectionSlotResult) *string { return v.SubnetResourceId }).(pulumi.StringPtrOutput)
}

// A flag that specifies if the scale unit this Web App is on supports Swift integration.
func (o LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput) SwiftSupported() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSwiftVirtualNetworkConnectionSlotResult) *bool { return v.SwiftSupported }).(pulumi.BoolPtrOutput)
}

// Resource type.
func (o LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSwiftVirtualNetworkConnectionSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppSwiftVirtualNetworkConnectionSlotResultOutput{})
}