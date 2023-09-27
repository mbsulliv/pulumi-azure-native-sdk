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

// Description for Lists the secrets for an existing static site.
// Azure REST API version: 2022-09-01.
func ListStaticSiteSecrets(ctx *pulumi.Context, args *ListStaticSiteSecretsArgs, opts ...pulumi.InvokeOption) (*ListStaticSiteSecretsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListStaticSiteSecretsResult
	err := ctx.Invoke("azure-native:web:listStaticSiteSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStaticSiteSecretsArgs struct {
	// Name of the static site.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// String dictionary resource.
type ListStaticSiteSecretsResult struct {
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

func ListStaticSiteSecretsOutput(ctx *pulumi.Context, args ListStaticSiteSecretsOutputArgs, opts ...pulumi.InvokeOption) ListStaticSiteSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStaticSiteSecretsResult, error) {
			args := v.(ListStaticSiteSecretsArgs)
			r, err := ListStaticSiteSecrets(ctx, &args, opts...)
			var s ListStaticSiteSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStaticSiteSecretsResultOutput)
}

type ListStaticSiteSecretsOutputArgs struct {
	// Name of the static site.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListStaticSiteSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteSecretsArgs)(nil)).Elem()
}

// String dictionary resource.
type ListStaticSiteSecretsResultOutput struct{ *pulumi.OutputState }

func (ListStaticSiteSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteSecretsResult)(nil)).Elem()
}

func (o ListStaticSiteSecretsResultOutput) ToListStaticSiteSecretsResultOutput() ListStaticSiteSecretsResultOutput {
	return o
}

func (o ListStaticSiteSecretsResultOutput) ToListStaticSiteSecretsResultOutputWithContext(ctx context.Context) ListStaticSiteSecretsResultOutput {
	return o
}

func (o ListStaticSiteSecretsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListStaticSiteSecretsResult] {
	return pulumix.Output[ListStaticSiteSecretsResult]{
		OutputState: o.OutputState,
	}
}

// Resource Id.
func (o ListStaticSiteSecretsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteSecretsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o ListStaticSiteSecretsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListStaticSiteSecretsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o ListStaticSiteSecretsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteSecretsResult) string { return v.Name }).(pulumi.StringOutput)
}

// Settings.
func (o ListStaticSiteSecretsResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListStaticSiteSecretsResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

// Resource type.
func (o ListStaticSiteSecretsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteSecretsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStaticSiteSecretsResultOutput{})
}
