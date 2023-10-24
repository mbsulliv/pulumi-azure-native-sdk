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

// Description for Gets the connection strings of an app.
// Azure REST API version: 2022-09-01.
//
// Other available API versions: 2020-10-01.
func ListWebAppConnectionStrings(ctx *pulumi.Context, args *ListWebAppConnectionStringsArgs, opts ...pulumi.InvokeOption) (*ListWebAppConnectionStringsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWebAppConnectionStringsResult
	err := ctx.Invoke("azure-native:web:listWebAppConnectionStrings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppConnectionStringsArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// String dictionary resource.
type ListWebAppConnectionStringsResult struct {
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

func ListWebAppConnectionStringsOutput(ctx *pulumi.Context, args ListWebAppConnectionStringsOutputArgs, opts ...pulumi.InvokeOption) ListWebAppConnectionStringsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppConnectionStringsResult, error) {
			args := v.(ListWebAppConnectionStringsArgs)
			r, err := ListWebAppConnectionStrings(ctx, &args, opts...)
			var s ListWebAppConnectionStringsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppConnectionStringsResultOutput)
}

type ListWebAppConnectionStringsOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppConnectionStringsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppConnectionStringsArgs)(nil)).Elem()
}

// String dictionary resource.
type ListWebAppConnectionStringsResultOutput struct{ *pulumi.OutputState }

func (ListWebAppConnectionStringsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppConnectionStringsResult)(nil)).Elem()
}

func (o ListWebAppConnectionStringsResultOutput) ToListWebAppConnectionStringsResultOutput() ListWebAppConnectionStringsResultOutput {
	return o
}

func (o ListWebAppConnectionStringsResultOutput) ToListWebAppConnectionStringsResultOutputWithContext(ctx context.Context) ListWebAppConnectionStringsResultOutput {
	return o
}

func (o ListWebAppConnectionStringsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListWebAppConnectionStringsResult] {
	return pulumix.Output[ListWebAppConnectionStringsResult]{
		OutputState: o.OutputState,
	}
}

// Resource Id.
func (o ListWebAppConnectionStringsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppConnectionStringsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o ListWebAppConnectionStringsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppConnectionStringsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o ListWebAppConnectionStringsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppConnectionStringsResult) string { return v.Name }).(pulumi.StringOutput)
}

// Connection strings.
func (o ListWebAppConnectionStringsResultOutput) Properties() ConnStringValueTypePairResponseMapOutput {
	return o.ApplyT(func(v ListWebAppConnectionStringsResult) map[string]ConnStringValueTypePairResponse {
		return v.Properties
	}).(ConnStringValueTypePairResponseMapOutput)
}

// Resource type.
func (o ListWebAppConnectionStringsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppConnectionStringsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppConnectionStringsResultOutput{})
}
