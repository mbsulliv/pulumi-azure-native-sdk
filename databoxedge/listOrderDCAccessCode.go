// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package databoxedge

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// DC Access code in the case of Self Managed Shipping.
// Azure REST API version: 2022-03-01.
func ListOrderDCAccessCode(ctx *pulumi.Context, args *ListOrderDCAccessCodeArgs, opts ...pulumi.InvokeOption) (*ListOrderDCAccessCodeResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListOrderDCAccessCodeResult
	err := ctx.Invoke("azure-native:databoxedge:listOrderDCAccessCode", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListOrderDCAccessCodeArgs struct {
	// The device name
	DeviceName string `pulumi:"deviceName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// DC Access code in the case of Self Managed Shipping.
type ListOrderDCAccessCodeResult struct {
	// DCAccess Code for the Self Managed shipment.
	AuthCode *string `pulumi:"authCode"`
}

func ListOrderDCAccessCodeOutput(ctx *pulumi.Context, args ListOrderDCAccessCodeOutputArgs, opts ...pulumi.InvokeOption) ListOrderDCAccessCodeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListOrderDCAccessCodeResult, error) {
			args := v.(ListOrderDCAccessCodeArgs)
			r, err := ListOrderDCAccessCode(ctx, &args, opts...)
			var s ListOrderDCAccessCodeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListOrderDCAccessCodeResultOutput)
}

type ListOrderDCAccessCodeOutputArgs struct {
	// The device name
	DeviceName pulumi.StringInput `pulumi:"deviceName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListOrderDCAccessCodeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOrderDCAccessCodeArgs)(nil)).Elem()
}

// DC Access code in the case of Self Managed Shipping.
type ListOrderDCAccessCodeResultOutput struct{ *pulumi.OutputState }

func (ListOrderDCAccessCodeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOrderDCAccessCodeResult)(nil)).Elem()
}

func (o ListOrderDCAccessCodeResultOutput) ToListOrderDCAccessCodeResultOutput() ListOrderDCAccessCodeResultOutput {
	return o
}

func (o ListOrderDCAccessCodeResultOutput) ToListOrderDCAccessCodeResultOutputWithContext(ctx context.Context) ListOrderDCAccessCodeResultOutput {
	return o
}

func (o ListOrderDCAccessCodeResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListOrderDCAccessCodeResult] {
	return pulumix.Output[ListOrderDCAccessCodeResult]{
		OutputState: o.OutputState,
	}
}

// DCAccess Code for the Self Managed shipment.
func (o ListOrderDCAccessCodeResultOutput) AuthCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListOrderDCAccessCodeResult) *string { return v.AuthCode }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListOrderDCAccessCodeResultOutput{})
}
