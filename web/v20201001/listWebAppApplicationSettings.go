// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the application settings of an app.
func ListWebAppApplicationSettings(ctx *pulumi.Context, args *ListWebAppApplicationSettingsArgs, opts ...pulumi.InvokeOption) (*ListWebAppApplicationSettingsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWebAppApplicationSettingsResult
	err := ctx.Invoke("azure-native:web/v20201001:listWebAppApplicationSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppApplicationSettingsArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// String dictionary resource.
type ListWebAppApplicationSettingsResult struct {
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Settings.
	Properties map[string]string `pulumi:"properties"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

func ListWebAppApplicationSettingsOutput(ctx *pulumi.Context, args ListWebAppApplicationSettingsOutputArgs, opts ...pulumi.InvokeOption) ListWebAppApplicationSettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppApplicationSettingsResult, error) {
			args := v.(ListWebAppApplicationSettingsArgs)
			r, err := ListWebAppApplicationSettings(ctx, &args, opts...)
			var s ListWebAppApplicationSettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppApplicationSettingsResultOutput)
}

type ListWebAppApplicationSettingsOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppApplicationSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppApplicationSettingsArgs)(nil)).Elem()
}

// String dictionary resource.
type ListWebAppApplicationSettingsResultOutput struct{ *pulumi.OutputState }

func (ListWebAppApplicationSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppApplicationSettingsResult)(nil)).Elem()
}

func (o ListWebAppApplicationSettingsResultOutput) ToListWebAppApplicationSettingsResultOutput() ListWebAppApplicationSettingsResultOutput {
	return o
}

func (o ListWebAppApplicationSettingsResultOutput) ToListWebAppApplicationSettingsResultOutputWithContext(ctx context.Context) ListWebAppApplicationSettingsResultOutput {
	return o
}

func (o ListWebAppApplicationSettingsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListWebAppApplicationSettingsResult] {
	return pulumix.Output[ListWebAppApplicationSettingsResult]{
		OutputState: o.OutputState,
	}
}

// Resource Id.
func (o ListWebAppApplicationSettingsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppApplicationSettingsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o ListWebAppApplicationSettingsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppApplicationSettingsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o ListWebAppApplicationSettingsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppApplicationSettingsResult) string { return v.Name }).(pulumi.StringOutput)
}

// Settings.
func (o ListWebAppApplicationSettingsResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListWebAppApplicationSettingsResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

// The system metadata relating to this resource.
func (o ListWebAppApplicationSettingsResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v ListWebAppApplicationSettingsResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o ListWebAppApplicationSettingsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppApplicationSettingsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppApplicationSettingsResultOutput{})
}
