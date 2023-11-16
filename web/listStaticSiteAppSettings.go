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
// Other available API versions: 2021-02-01, 2023-01-01.
func ListStaticSiteAppSettings(ctx *pulumi.Context, args *ListStaticSiteAppSettingsArgs, opts ...pulumi.InvokeOption) (*ListStaticSiteAppSettingsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListStaticSiteAppSettingsResult
	err := ctx.Invoke("azure-native:web:listStaticSiteAppSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStaticSiteAppSettingsArgs struct {
	// Name of the static site.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// String dictionary resource.
type ListStaticSiteAppSettingsResult struct {
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

func ListStaticSiteAppSettingsOutput(ctx *pulumi.Context, args ListStaticSiteAppSettingsOutputArgs, opts ...pulumi.InvokeOption) ListStaticSiteAppSettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStaticSiteAppSettingsResult, error) {
			args := v.(ListStaticSiteAppSettingsArgs)
			r, err := ListStaticSiteAppSettings(ctx, &args, opts...)
			var s ListStaticSiteAppSettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStaticSiteAppSettingsResultOutput)
}

type ListStaticSiteAppSettingsOutputArgs struct {
	// Name of the static site.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListStaticSiteAppSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteAppSettingsArgs)(nil)).Elem()
}

// String dictionary resource.
type ListStaticSiteAppSettingsResultOutput struct{ *pulumi.OutputState }

func (ListStaticSiteAppSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteAppSettingsResult)(nil)).Elem()
}

func (o ListStaticSiteAppSettingsResultOutput) ToListStaticSiteAppSettingsResultOutput() ListStaticSiteAppSettingsResultOutput {
	return o
}

func (o ListStaticSiteAppSettingsResultOutput) ToListStaticSiteAppSettingsResultOutputWithContext(ctx context.Context) ListStaticSiteAppSettingsResultOutput {
	return o
}

// Resource Id.
func (o ListStaticSiteAppSettingsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteAppSettingsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o ListStaticSiteAppSettingsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListStaticSiteAppSettingsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o ListStaticSiteAppSettingsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteAppSettingsResult) string { return v.Name }).(pulumi.StringOutput)
}

// Settings.
func (o ListStaticSiteAppSettingsResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListStaticSiteAppSettingsResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

// Resource type.
func (o ListStaticSiteAppSettingsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteAppSettingsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStaticSiteAppSettingsResultOutput{})
}
