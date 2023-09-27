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

// Description for Gets the application settings of a static site build.
func ListStaticSiteBuildAppSettings(ctx *pulumi.Context, args *ListStaticSiteBuildAppSettingsArgs, opts ...pulumi.InvokeOption) (*ListStaticSiteBuildAppSettingsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListStaticSiteBuildAppSettingsResult
	err := ctx.Invoke("azure-native:web/v20220901:listStaticSiteBuildAppSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStaticSiteBuildAppSettingsArgs struct {
	// The stage site identifier.
	EnvironmentName string `pulumi:"environmentName"`
	// Name of the static site.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// String dictionary resource.
type ListStaticSiteBuildAppSettingsResult struct {
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

func ListStaticSiteBuildAppSettingsOutput(ctx *pulumi.Context, args ListStaticSiteBuildAppSettingsOutputArgs, opts ...pulumi.InvokeOption) ListStaticSiteBuildAppSettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStaticSiteBuildAppSettingsResult, error) {
			args := v.(ListStaticSiteBuildAppSettingsArgs)
			r, err := ListStaticSiteBuildAppSettings(ctx, &args, opts...)
			var s ListStaticSiteBuildAppSettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStaticSiteBuildAppSettingsResultOutput)
}

type ListStaticSiteBuildAppSettingsOutputArgs struct {
	// The stage site identifier.
	EnvironmentName pulumi.StringInput `pulumi:"environmentName"`
	// Name of the static site.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListStaticSiteBuildAppSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteBuildAppSettingsArgs)(nil)).Elem()
}

// String dictionary resource.
type ListStaticSiteBuildAppSettingsResultOutput struct{ *pulumi.OutputState }

func (ListStaticSiteBuildAppSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteBuildAppSettingsResult)(nil)).Elem()
}

func (o ListStaticSiteBuildAppSettingsResultOutput) ToListStaticSiteBuildAppSettingsResultOutput() ListStaticSiteBuildAppSettingsResultOutput {
	return o
}

func (o ListStaticSiteBuildAppSettingsResultOutput) ToListStaticSiteBuildAppSettingsResultOutputWithContext(ctx context.Context) ListStaticSiteBuildAppSettingsResultOutput {
	return o
}

func (o ListStaticSiteBuildAppSettingsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListStaticSiteBuildAppSettingsResult] {
	return pulumix.Output[ListStaticSiteBuildAppSettingsResult]{
		OutputState: o.OutputState,
	}
}

// Resource Id.
func (o ListStaticSiteBuildAppSettingsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteBuildAppSettingsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o ListStaticSiteBuildAppSettingsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListStaticSiteBuildAppSettingsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o ListStaticSiteBuildAppSettingsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteBuildAppSettingsResult) string { return v.Name }).(pulumi.StringOutput)
}

// Settings.
func (o ListStaticSiteBuildAppSettingsResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListStaticSiteBuildAppSettingsResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

// Resource type.
func (o ListStaticSiteBuildAppSettingsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteBuildAppSettingsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStaticSiteBuildAppSettingsResultOutput{})
}
