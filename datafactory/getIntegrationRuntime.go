// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets an integration runtime.
// Azure REST API version: 2018-06-01.
func LookupIntegrationRuntime(ctx *pulumi.Context, args *LookupIntegrationRuntimeArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationRuntimeResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIntegrationRuntimeResult
	err := ctx.Invoke("azure-native:datafactory:getIntegrationRuntime", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationRuntimeArgs struct {
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// The integration runtime name.
	IntegrationRuntimeName string `pulumi:"integrationRuntimeName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Integration runtime resource type.
type LookupIntegrationRuntimeResult struct {
	// Etag identifies change in the resource.
	Etag string `pulumi:"etag"`
	// The resource identifier.
	Id string `pulumi:"id"`
	// The resource name.
	Name string `pulumi:"name"`
	// Integration runtime properties.
	Properties interface{} `pulumi:"properties"`
	// The resource type.
	Type string `pulumi:"type"`
}

func LookupIntegrationRuntimeOutput(ctx *pulumi.Context, args LookupIntegrationRuntimeOutputArgs, opts ...pulumi.InvokeOption) LookupIntegrationRuntimeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIntegrationRuntimeResult, error) {
			args := v.(LookupIntegrationRuntimeArgs)
			r, err := LookupIntegrationRuntime(ctx, &args, opts...)
			var s LookupIntegrationRuntimeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIntegrationRuntimeResultOutput)
}

type LookupIntegrationRuntimeOutputArgs struct {
	// The factory name.
	FactoryName pulumi.StringInput `pulumi:"factoryName"`
	// The integration runtime name.
	IntegrationRuntimeName pulumi.StringInput `pulumi:"integrationRuntimeName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIntegrationRuntimeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationRuntimeArgs)(nil)).Elem()
}

// Integration runtime resource type.
type LookupIntegrationRuntimeResultOutput struct{ *pulumi.OutputState }

func (LookupIntegrationRuntimeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationRuntimeResult)(nil)).Elem()
}

func (o LookupIntegrationRuntimeResultOutput) ToLookupIntegrationRuntimeResultOutput() LookupIntegrationRuntimeResultOutput {
	return o
}

func (o LookupIntegrationRuntimeResultOutput) ToLookupIntegrationRuntimeResultOutputWithContext(ctx context.Context) LookupIntegrationRuntimeResultOutput {
	return o
}

func (o LookupIntegrationRuntimeResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupIntegrationRuntimeResult] {
	return pulumix.Output[LookupIntegrationRuntimeResult]{
		OutputState: o.OutputState,
	}
}

// Etag identifies change in the resource.
func (o LookupIntegrationRuntimeResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationRuntimeResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The resource identifier.
func (o LookupIntegrationRuntimeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationRuntimeResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource name.
func (o LookupIntegrationRuntimeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationRuntimeResult) string { return v.Name }).(pulumi.StringOutput)
}

// Integration runtime properties.
func (o LookupIntegrationRuntimeResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupIntegrationRuntimeResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

// The resource type.
func (o LookupIntegrationRuntimeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationRuntimeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntegrationRuntimeResultOutput{})
}
