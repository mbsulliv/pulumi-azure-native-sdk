// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containerregistry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the detailed information for a given pipeline run.
// Azure REST API version: 2023-01-01-preview.
//
// Other available API versions: 2023-06-01-preview, 2023-08-01-preview.
func LookupPipelineRun(ctx *pulumi.Context, args *LookupPipelineRunArgs, opts ...pulumi.InvokeOption) (*LookupPipelineRunResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPipelineRunResult
	err := ctx.Invoke("azure-native:containerregistry:getPipelineRun", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPipelineRunArgs struct {
	// The name of the pipeline run.
	PipelineRunName string `pulumi:"pipelineRunName"`
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An object that represents a pipeline run for a container registry.
type LookupPipelineRunResult struct {
	// How the pipeline run should be forced to recreate even if the pipeline run configuration has not changed.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// The resource ID.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The provisioning state of a pipeline run.
	ProvisioningState string `pulumi:"provisioningState"`
	// The request parameters for a pipeline run.
	Request *PipelineRunRequestResponse `pulumi:"request"`
	// The response of a pipeline run.
	Response PipelineRunResponseResponse `pulumi:"response"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupPipelineRunResult
func (val *LookupPipelineRunResult) Defaults() *LookupPipelineRunResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Request = tmp.Request.Defaults()

	tmp.Response = *tmp.Response.Defaults()

	return &tmp
}

func LookupPipelineRunOutput(ctx *pulumi.Context, args LookupPipelineRunOutputArgs, opts ...pulumi.InvokeOption) LookupPipelineRunResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPipelineRunResult, error) {
			args := v.(LookupPipelineRunArgs)
			r, err := LookupPipelineRun(ctx, &args, opts...)
			var s LookupPipelineRunResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPipelineRunResultOutput)
}

type LookupPipelineRunOutputArgs struct {
	// The name of the pipeline run.
	PipelineRunName pulumi.StringInput `pulumi:"pipelineRunName"`
	// The name of the container registry.
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPipelineRunOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineRunArgs)(nil)).Elem()
}

// An object that represents a pipeline run for a container registry.
type LookupPipelineRunResultOutput struct{ *pulumi.OutputState }

func (LookupPipelineRunResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineRunResult)(nil)).Elem()
}

func (o LookupPipelineRunResultOutput) ToLookupPipelineRunResultOutput() LookupPipelineRunResultOutput {
	return o
}

func (o LookupPipelineRunResultOutput) ToLookupPipelineRunResultOutputWithContext(ctx context.Context) LookupPipelineRunResultOutput {
	return o
}

func (o LookupPipelineRunResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPipelineRunResult] {
	return pulumix.Output[LookupPipelineRunResult]{
		OutputState: o.OutputState,
	}
}

// How the pipeline run should be forced to recreate even if the pipeline run configuration has not changed.
func (o LookupPipelineRunResultOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPipelineRunResult) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

// The resource ID.
func (o LookupPipelineRunResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineRunResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupPipelineRunResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineRunResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of a pipeline run.
func (o LookupPipelineRunResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineRunResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The request parameters for a pipeline run.
func (o LookupPipelineRunResultOutput) Request() PipelineRunRequestResponsePtrOutput {
	return o.ApplyT(func(v LookupPipelineRunResult) *PipelineRunRequestResponse { return v.Request }).(PipelineRunRequestResponsePtrOutput)
}

// The response of a pipeline run.
func (o LookupPipelineRunResultOutput) Response() PipelineRunResponseResponseOutput {
	return o.ApplyT(func(v LookupPipelineRunResult) PipelineRunResponseResponse { return v.Response }).(PipelineRunResponseResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupPipelineRunResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPipelineRunResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupPipelineRunResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineRunResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPipelineRunResultOutput{})
}
