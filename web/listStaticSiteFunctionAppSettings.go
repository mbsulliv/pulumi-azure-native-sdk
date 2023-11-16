// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description for Gets the application settings of a static site.
// Azure REST API version: 2022-09-01.
//
// Other available API versions: 2020-10-01, 2021-02-01, 2023-01-01.
func ListStaticSiteFunctionAppSettings(ctx *pulumi.Context, args *ListStaticSiteFunctionAppSettingsArgs, opts ...pulumi.InvokeOption) (*ListStaticSiteFunctionAppSettingsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListStaticSiteFunctionAppSettingsResult
	err := ctx.Invoke("azure-native:web:listStaticSiteFunctionAppSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStaticSiteFunctionAppSettingsArgs struct {
	// Name of the static site.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// String dictionary resource.
type ListStaticSiteFunctionAppSettingsResult struct {
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

func ListStaticSiteFunctionAppSettingsOutput(ctx *pulumi.Context, args ListStaticSiteFunctionAppSettingsOutputArgs, opts ...pulumi.InvokeOption) ListStaticSiteFunctionAppSettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStaticSiteFunctionAppSettingsResult, error) {
			args := v.(ListStaticSiteFunctionAppSettingsArgs)
			r, err := ListStaticSiteFunctionAppSettings(ctx, &args, opts...)
			var s ListStaticSiteFunctionAppSettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStaticSiteFunctionAppSettingsResultOutput)
}

type ListStaticSiteFunctionAppSettingsOutputArgs struct {
	// Name of the static site.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListStaticSiteFunctionAppSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteFunctionAppSettingsArgs)(nil)).Elem()
}

// String dictionary resource.
type ListStaticSiteFunctionAppSettingsResultOutput struct{ *pulumi.OutputState }

func (ListStaticSiteFunctionAppSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteFunctionAppSettingsResult)(nil)).Elem()
}

func (o ListStaticSiteFunctionAppSettingsResultOutput) ToListStaticSiteFunctionAppSettingsResultOutput() ListStaticSiteFunctionAppSettingsResultOutput {
	return o
}

func (o ListStaticSiteFunctionAppSettingsResultOutput) ToListStaticSiteFunctionAppSettingsResultOutputWithContext(ctx context.Context) ListStaticSiteFunctionAppSettingsResultOutput {
	return o
}

// Resource Id.
func (o ListStaticSiteFunctionAppSettingsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteFunctionAppSettingsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o ListStaticSiteFunctionAppSettingsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListStaticSiteFunctionAppSettingsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o ListStaticSiteFunctionAppSettingsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteFunctionAppSettingsResult) string { return v.Name }).(pulumi.StringOutput)
}

// Settings.
func (o ListStaticSiteFunctionAppSettingsResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListStaticSiteFunctionAppSettingsResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

// Resource type.
func (o ListStaticSiteFunctionAppSettingsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteFunctionAppSettingsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStaticSiteFunctionAppSettingsResultOutput{})
}
