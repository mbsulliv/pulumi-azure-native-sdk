// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package migrate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a Hyper-V collector.
// Azure REST API version: 2019-10-01.
func LookupHyperVCollector(ctx *pulumi.Context, args *LookupHyperVCollectorArgs, opts ...pulumi.InvokeOption) (*LookupHyperVCollectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupHyperVCollectorResult
	err := ctx.Invoke("azure-native:migrate:getHyperVCollector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHyperVCollectorArgs struct {
	// Unique name of a Hyper-V collector within a project.
	HyperVCollectorName string `pulumi:"hyperVCollectorName"`
	// Name of the Azure Migrate project.
	ProjectName string `pulumi:"projectName"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupHyperVCollectorResult struct {
	ETag       *string                     `pulumi:"eTag"`
	Id         string                      `pulumi:"id"`
	Name       string                      `pulumi:"name"`
	Properties CollectorPropertiesResponse `pulumi:"properties"`
	Type       string                      `pulumi:"type"`
}

func LookupHyperVCollectorOutput(ctx *pulumi.Context, args LookupHyperVCollectorOutputArgs, opts ...pulumi.InvokeOption) LookupHyperVCollectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHyperVCollectorResult, error) {
			args := v.(LookupHyperVCollectorArgs)
			r, err := LookupHyperVCollector(ctx, &args, opts...)
			var s LookupHyperVCollectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHyperVCollectorResultOutput)
}

type LookupHyperVCollectorOutputArgs struct {
	// Unique name of a Hyper-V collector within a project.
	HyperVCollectorName pulumi.StringInput `pulumi:"hyperVCollectorName"`
	// Name of the Azure Migrate project.
	ProjectName pulumi.StringInput `pulumi:"projectName"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupHyperVCollectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHyperVCollectorArgs)(nil)).Elem()
}

type LookupHyperVCollectorResultOutput struct{ *pulumi.OutputState }

func (LookupHyperVCollectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHyperVCollectorResult)(nil)).Elem()
}

func (o LookupHyperVCollectorResultOutput) ToLookupHyperVCollectorResultOutput() LookupHyperVCollectorResultOutput {
	return o
}

func (o LookupHyperVCollectorResultOutput) ToLookupHyperVCollectorResultOutputWithContext(ctx context.Context) LookupHyperVCollectorResultOutput {
	return o
}

func (o LookupHyperVCollectorResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHyperVCollectorResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupHyperVCollectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHyperVCollectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupHyperVCollectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHyperVCollectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupHyperVCollectorResultOutput) Properties() CollectorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupHyperVCollectorResult) CollectorPropertiesResponse { return v.Properties }).(CollectorPropertiesResponseOutput)
}

func (o LookupHyperVCollectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHyperVCollectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHyperVCollectorResultOutput{})
}
