// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearning

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve an Azure ML commitment plan by its subscription, resource group and name.
// Azure REST API version: 2016-05-01-preview.
func LookupCommitmentPlan(ctx *pulumi.Context, args *LookupCommitmentPlanArgs, opts ...pulumi.InvokeOption) (*LookupCommitmentPlanResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCommitmentPlanResult
	err := ctx.Invoke("azure-native:machinelearning:getCommitmentPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCommitmentPlanArgs struct {
	// The Azure ML commitment plan name.
	CommitmentPlanName string `pulumi:"commitmentPlanName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Azure ML commitment plan resource.
type LookupCommitmentPlanResult struct {
	// An entity tag used to enforce optimistic concurrency.
	Etag *string `pulumi:"etag"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The commitment plan properties.
	Properties CommitmentPlanPropertiesResponse `pulumi:"properties"`
	// The commitment plan SKU.
	Sku *ResourceSkuResponse `pulumi:"sku"`
	// User-defined tags for the resource.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupCommitmentPlanOutput(ctx *pulumi.Context, args LookupCommitmentPlanOutputArgs, opts ...pulumi.InvokeOption) LookupCommitmentPlanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCommitmentPlanResult, error) {
			args := v.(LookupCommitmentPlanArgs)
			r, err := LookupCommitmentPlan(ctx, &args, opts...)
			var s LookupCommitmentPlanResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCommitmentPlanResultOutput)
}

type LookupCommitmentPlanOutputArgs struct {
	// The Azure ML commitment plan name.
	CommitmentPlanName pulumi.StringInput `pulumi:"commitmentPlanName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCommitmentPlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCommitmentPlanArgs)(nil)).Elem()
}

// An Azure ML commitment plan resource.
type LookupCommitmentPlanResultOutput struct{ *pulumi.OutputState }

func (LookupCommitmentPlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCommitmentPlanResult)(nil)).Elem()
}

func (o LookupCommitmentPlanResultOutput) ToLookupCommitmentPlanResultOutput() LookupCommitmentPlanResultOutput {
	return o
}

func (o LookupCommitmentPlanResultOutput) ToLookupCommitmentPlanResultOutputWithContext(ctx context.Context) LookupCommitmentPlanResultOutput {
	return o
}

// An entity tag used to enforce optimistic concurrency.
func (o LookupCommitmentPlanResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCommitmentPlanResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupCommitmentPlanResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommitmentPlanResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupCommitmentPlanResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommitmentPlanResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupCommitmentPlanResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommitmentPlanResult) string { return v.Name }).(pulumi.StringOutput)
}

// The commitment plan properties.
func (o LookupCommitmentPlanResultOutput) Properties() CommitmentPlanPropertiesResponseOutput {
	return o.ApplyT(func(v LookupCommitmentPlanResult) CommitmentPlanPropertiesResponse { return v.Properties }).(CommitmentPlanPropertiesResponseOutput)
}

// The commitment plan SKU.
func (o LookupCommitmentPlanResultOutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupCommitmentPlanResult) *ResourceSkuResponse { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

// User-defined tags for the resource.
func (o LookupCommitmentPlanResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCommitmentPlanResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupCommitmentPlanResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommitmentPlanResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCommitmentPlanResultOutput{})
}
