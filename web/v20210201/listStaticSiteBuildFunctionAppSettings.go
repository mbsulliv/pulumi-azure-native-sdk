// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the application settings of a static site build.
func ListStaticSiteBuildFunctionAppSettings(ctx *pulumi.Context, args *ListStaticSiteBuildFunctionAppSettingsArgs, opts ...pulumi.InvokeOption) (*ListStaticSiteBuildFunctionAppSettingsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListStaticSiteBuildFunctionAppSettingsResult
	err := ctx.Invoke("azure-native:web/v20210201:listStaticSiteBuildFunctionAppSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStaticSiteBuildFunctionAppSettingsArgs struct {
	// The stage site identifier.
	EnvironmentName string `pulumi:"environmentName"`
	// Name of the static site.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// String dictionary resource.
type ListStaticSiteBuildFunctionAppSettingsResult struct {
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

func ListStaticSiteBuildFunctionAppSettingsOutput(ctx *pulumi.Context, args ListStaticSiteBuildFunctionAppSettingsOutputArgs, opts ...pulumi.InvokeOption) ListStaticSiteBuildFunctionAppSettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStaticSiteBuildFunctionAppSettingsResult, error) {
			args := v.(ListStaticSiteBuildFunctionAppSettingsArgs)
			r, err := ListStaticSiteBuildFunctionAppSettings(ctx, &args, opts...)
			var s ListStaticSiteBuildFunctionAppSettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStaticSiteBuildFunctionAppSettingsResultOutput)
}

type ListStaticSiteBuildFunctionAppSettingsOutputArgs struct {
	// The stage site identifier.
	EnvironmentName pulumi.StringInput `pulumi:"environmentName"`
	// Name of the static site.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListStaticSiteBuildFunctionAppSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteBuildFunctionAppSettingsArgs)(nil)).Elem()
}

// String dictionary resource.
type ListStaticSiteBuildFunctionAppSettingsResultOutput struct{ *pulumi.OutputState }

func (ListStaticSiteBuildFunctionAppSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteBuildFunctionAppSettingsResult)(nil)).Elem()
}

func (o ListStaticSiteBuildFunctionAppSettingsResultOutput) ToListStaticSiteBuildFunctionAppSettingsResultOutput() ListStaticSiteBuildFunctionAppSettingsResultOutput {
	return o
}

func (o ListStaticSiteBuildFunctionAppSettingsResultOutput) ToListStaticSiteBuildFunctionAppSettingsResultOutputWithContext(ctx context.Context) ListStaticSiteBuildFunctionAppSettingsResultOutput {
	return o
}

// Resource Id.
func (o ListStaticSiteBuildFunctionAppSettingsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteBuildFunctionAppSettingsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o ListStaticSiteBuildFunctionAppSettingsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListStaticSiteBuildFunctionAppSettingsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o ListStaticSiteBuildFunctionAppSettingsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteBuildFunctionAppSettingsResult) string { return v.Name }).(pulumi.StringOutput)
}

// Settings.
func (o ListStaticSiteBuildFunctionAppSettingsResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListStaticSiteBuildFunctionAppSettingsResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

// Resource type.
func (o ListStaticSiteBuildFunctionAppSettingsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteBuildFunctionAppSettingsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStaticSiteBuildFunctionAppSettingsResultOutput{})
}
