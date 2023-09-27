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

// Gets a linked service.
// Azure REST API version: 2018-06-01.
func LookupLinkedService(ctx *pulumi.Context, args *LookupLinkedServiceArgs, opts ...pulumi.InvokeOption) (*LookupLinkedServiceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupLinkedServiceResult
	err := ctx.Invoke("azure-native:datafactory:getLinkedService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkedServiceArgs struct {
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// The linked service name.
	LinkedServiceName string `pulumi:"linkedServiceName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Linked service resource type.
type LookupLinkedServiceResult struct {
	// Etag identifies change in the resource.
	Etag string `pulumi:"etag"`
	// The resource identifier.
	Id string `pulumi:"id"`
	// The resource name.
	Name string `pulumi:"name"`
	// Properties of linked service.
	Properties interface{} `pulumi:"properties"`
	// The resource type.
	Type string `pulumi:"type"`
}

func LookupLinkedServiceOutput(ctx *pulumi.Context, args LookupLinkedServiceOutputArgs, opts ...pulumi.InvokeOption) LookupLinkedServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLinkedServiceResult, error) {
			args := v.(LookupLinkedServiceArgs)
			r, err := LookupLinkedService(ctx, &args, opts...)
			var s LookupLinkedServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLinkedServiceResultOutput)
}

type LookupLinkedServiceOutputArgs struct {
	// The factory name.
	FactoryName pulumi.StringInput `pulumi:"factoryName"`
	// The linked service name.
	LinkedServiceName pulumi.StringInput `pulumi:"linkedServiceName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLinkedServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkedServiceArgs)(nil)).Elem()
}

// Linked service resource type.
type LookupLinkedServiceResultOutput struct{ *pulumi.OutputState }

func (LookupLinkedServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkedServiceResult)(nil)).Elem()
}

func (o LookupLinkedServiceResultOutput) ToLookupLinkedServiceResultOutput() LookupLinkedServiceResultOutput {
	return o
}

func (o LookupLinkedServiceResultOutput) ToLookupLinkedServiceResultOutputWithContext(ctx context.Context) LookupLinkedServiceResultOutput {
	return o
}

func (o LookupLinkedServiceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupLinkedServiceResult] {
	return pulumix.Output[LookupLinkedServiceResult]{
		OutputState: o.OutputState,
	}
}

// Etag identifies change in the resource.
func (o LookupLinkedServiceResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The resource identifier.
func (o LookupLinkedServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource name.
func (o LookupLinkedServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of linked service.
func (o LookupLinkedServiceResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

// The resource type.
func (o LookupLinkedServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLinkedServiceResultOutput{})
}
