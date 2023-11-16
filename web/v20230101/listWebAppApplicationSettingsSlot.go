// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description for Gets the application settings of an app.
func ListWebAppApplicationSettingsSlot(ctx *pulumi.Context, args *ListWebAppApplicationSettingsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppApplicationSettingsSlotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWebAppApplicationSettingsSlotResult
	err := ctx.Invoke("azure-native:web/v20230101:listWebAppApplicationSettingsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppApplicationSettingsSlotArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will get the application settings for the production slot.
	Slot string `pulumi:"slot"`
}

// String dictionary resource.
type ListWebAppApplicationSettingsSlotResult struct {
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Settings.
	Properties map[string]string `pulumi:"properties"`
	// Resource type.
	Type string `pulumi:"type"`
}

func ListWebAppApplicationSettingsSlotOutput(ctx *pulumi.Context, args ListWebAppApplicationSettingsSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppApplicationSettingsSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppApplicationSettingsSlotResult, error) {
			args := v.(ListWebAppApplicationSettingsSlotArgs)
			r, err := ListWebAppApplicationSettingsSlot(ctx, &args, opts...)
			var s ListWebAppApplicationSettingsSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppApplicationSettingsSlotResultOutput)
}

type ListWebAppApplicationSettingsSlotOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will get the application settings for the production slot.
	Slot pulumi.StringInput `pulumi:"slot"`
}

func (ListWebAppApplicationSettingsSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppApplicationSettingsSlotArgs)(nil)).Elem()
}

// String dictionary resource.
type ListWebAppApplicationSettingsSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppApplicationSettingsSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppApplicationSettingsSlotResult)(nil)).Elem()
}

func (o ListWebAppApplicationSettingsSlotResultOutput) ToListWebAppApplicationSettingsSlotResultOutput() ListWebAppApplicationSettingsSlotResultOutput {
	return o
}

func (o ListWebAppApplicationSettingsSlotResultOutput) ToListWebAppApplicationSettingsSlotResultOutputWithContext(ctx context.Context) ListWebAppApplicationSettingsSlotResultOutput {
	return o
}

// Resource Id.
func (o ListWebAppApplicationSettingsSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppApplicationSettingsSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o ListWebAppApplicationSettingsSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppApplicationSettingsSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o ListWebAppApplicationSettingsSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppApplicationSettingsSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

// Settings.
func (o ListWebAppApplicationSettingsSlotResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListWebAppApplicationSettingsSlotResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

// Resource type.
func (o ListWebAppApplicationSettingsSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppApplicationSettingsSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppApplicationSettingsSlotResultOutput{})
}
