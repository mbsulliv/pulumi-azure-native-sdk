// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Description for Gets the connection strings of an app.
func ListWebAppConnectionStringsSlot(ctx *pulumi.Context, args *ListWebAppConnectionStringsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppConnectionStringsSlotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWebAppConnectionStringsSlotResult
	err := ctx.Invoke("azure-native:web/v20220901:listWebAppConnectionStringsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppConnectionStringsSlotArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will get the connection settings for the production slot.
	Slot string `pulumi:"slot"`
}

// String dictionary resource.
type ListWebAppConnectionStringsSlotResult struct {
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Connection strings.
	Properties map[string]ConnStringValueTypePairResponse `pulumi:"properties"`
	// Resource type.
	Type string `pulumi:"type"`
}

func ListWebAppConnectionStringsSlotOutput(ctx *pulumi.Context, args ListWebAppConnectionStringsSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppConnectionStringsSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppConnectionStringsSlotResult, error) {
			args := v.(ListWebAppConnectionStringsSlotArgs)
			r, err := ListWebAppConnectionStringsSlot(ctx, &args, opts...)
			var s ListWebAppConnectionStringsSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppConnectionStringsSlotResultOutput)
}

type ListWebAppConnectionStringsSlotOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will get the connection settings for the production slot.
	Slot pulumi.StringInput `pulumi:"slot"`
}

func (ListWebAppConnectionStringsSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppConnectionStringsSlotArgs)(nil)).Elem()
}

// String dictionary resource.
type ListWebAppConnectionStringsSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppConnectionStringsSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppConnectionStringsSlotResult)(nil)).Elem()
}

func (o ListWebAppConnectionStringsSlotResultOutput) ToListWebAppConnectionStringsSlotResultOutput() ListWebAppConnectionStringsSlotResultOutput {
	return o
}

func (o ListWebAppConnectionStringsSlotResultOutput) ToListWebAppConnectionStringsSlotResultOutputWithContext(ctx context.Context) ListWebAppConnectionStringsSlotResultOutput {
	return o
}

func (o ListWebAppConnectionStringsSlotResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListWebAppConnectionStringsSlotResult] {
	return pulumix.Output[ListWebAppConnectionStringsSlotResult]{
		OutputState: o.OutputState,
	}
}

// Resource Id.
func (o ListWebAppConnectionStringsSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppConnectionStringsSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o ListWebAppConnectionStringsSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppConnectionStringsSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o ListWebAppConnectionStringsSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppConnectionStringsSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

// Connection strings.
func (o ListWebAppConnectionStringsSlotResultOutput) Properties() ConnStringValueTypePairResponseMapOutput {
	return o.ApplyT(func(v ListWebAppConnectionStringsSlotResult) map[string]ConnStringValueTypePairResponse {
		return v.Properties
	}).(ConnStringValueTypePairResponseMapOutput)
}

// Resource type.
func (o ListWebAppConnectionStringsSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppConnectionStringsSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppConnectionStringsSlotResultOutput{})
}
