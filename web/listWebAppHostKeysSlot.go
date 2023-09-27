// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Description for Get host secrets for a function app.
// Azure REST API version: 2022-09-01.
func ListWebAppHostKeysSlot(ctx *pulumi.Context, args *ListWebAppHostKeysSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppHostKeysSlotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWebAppHostKeysSlotResult
	err := ctx.Invoke("azure-native:web:listWebAppHostKeysSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppHostKeysSlotArgs struct {
	// Site name.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot.
	Slot string `pulumi:"slot"`
}

// Functions host level keys.
type ListWebAppHostKeysSlotResult struct {
	// Host level function keys.
	FunctionKeys map[string]string `pulumi:"functionKeys"`
	// Secret key.
	MasterKey *string `pulumi:"masterKey"`
	// System keys.
	SystemKeys map[string]string `pulumi:"systemKeys"`
}

func ListWebAppHostKeysSlotOutput(ctx *pulumi.Context, args ListWebAppHostKeysSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppHostKeysSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppHostKeysSlotResult, error) {
			args := v.(ListWebAppHostKeysSlotArgs)
			r, err := ListWebAppHostKeysSlot(ctx, &args, opts...)
			var s ListWebAppHostKeysSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppHostKeysSlotResultOutput)
}

type ListWebAppHostKeysSlotOutputArgs struct {
	// Site name.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the deployment slot.
	Slot pulumi.StringInput `pulumi:"slot"`
}

func (ListWebAppHostKeysSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppHostKeysSlotArgs)(nil)).Elem()
}

// Functions host level keys.
type ListWebAppHostKeysSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppHostKeysSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppHostKeysSlotResult)(nil)).Elem()
}

func (o ListWebAppHostKeysSlotResultOutput) ToListWebAppHostKeysSlotResultOutput() ListWebAppHostKeysSlotResultOutput {
	return o
}

func (o ListWebAppHostKeysSlotResultOutput) ToListWebAppHostKeysSlotResultOutputWithContext(ctx context.Context) ListWebAppHostKeysSlotResultOutput {
	return o
}

func (o ListWebAppHostKeysSlotResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListWebAppHostKeysSlotResult] {
	return pulumix.Output[ListWebAppHostKeysSlotResult]{
		OutputState: o.OutputState,
	}
}

// Host level function keys.
func (o ListWebAppHostKeysSlotResultOutput) FunctionKeys() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListWebAppHostKeysSlotResult) map[string]string { return v.FunctionKeys }).(pulumi.StringMapOutput)
}

// Secret key.
func (o ListWebAppHostKeysSlotResultOutput) MasterKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppHostKeysSlotResult) *string { return v.MasterKey }).(pulumi.StringPtrOutput)
}

// System keys.
func (o ListWebAppHostKeysSlotResultOutput) SystemKeys() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListWebAppHostKeysSlotResult) map[string]string { return v.SystemKeys }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppHostKeysSlotResultOutput{})
}
