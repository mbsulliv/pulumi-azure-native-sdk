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

// Gets the properties of the import pipeline.
// Azure REST API version: 2023-01-01-preview.
//
// Other available API versions: 2023-06-01-preview, 2023-08-01-preview.
func LookupImportPipeline(ctx *pulumi.Context, args *LookupImportPipelineArgs, opts ...pulumi.InvokeOption) (*LookupImportPipelineResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupImportPipelineResult
	err := ctx.Invoke("azure-native:containerregistry:getImportPipeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupImportPipelineArgs struct {
	// The name of the import pipeline.
	ImportPipelineName string `pulumi:"importPipelineName"`
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An object that represents an import pipeline for a container registry.
type LookupImportPipelineResult struct {
	// The resource ID.
	Id string `pulumi:"id"`
	// The identity of the import pipeline.
	Identity *IdentityPropertiesResponse `pulumi:"identity"`
	// The location of the import pipeline.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The list of all options configured for the pipeline.
	Options []string `pulumi:"options"`
	// The provisioning state of the pipeline at the time the operation was called.
	ProvisioningState string `pulumi:"provisioningState"`
	// The source properties of the import pipeline.
	Source ImportPipelineSourcePropertiesResponse `pulumi:"source"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The properties that describe the trigger of the import pipeline.
	Trigger *PipelineTriggerPropertiesResponse `pulumi:"trigger"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupImportPipelineResult
func (val *LookupImportPipelineResult) Defaults() *LookupImportPipelineResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Source = *tmp.Source.Defaults()

	tmp.Trigger = tmp.Trigger.Defaults()

	return &tmp
}

func LookupImportPipelineOutput(ctx *pulumi.Context, args LookupImportPipelineOutputArgs, opts ...pulumi.InvokeOption) LookupImportPipelineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupImportPipelineResult, error) {
			args := v.(LookupImportPipelineArgs)
			r, err := LookupImportPipeline(ctx, &args, opts...)
			var s LookupImportPipelineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupImportPipelineResultOutput)
}

type LookupImportPipelineOutputArgs struct {
	// The name of the import pipeline.
	ImportPipelineName pulumi.StringInput `pulumi:"importPipelineName"`
	// The name of the container registry.
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupImportPipelineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImportPipelineArgs)(nil)).Elem()
}

// An object that represents an import pipeline for a container registry.
type LookupImportPipelineResultOutput struct{ *pulumi.OutputState }

func (LookupImportPipelineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImportPipelineResult)(nil)).Elem()
}

func (o LookupImportPipelineResultOutput) ToLookupImportPipelineResultOutput() LookupImportPipelineResultOutput {
	return o
}

func (o LookupImportPipelineResultOutput) ToLookupImportPipelineResultOutputWithContext(ctx context.Context) LookupImportPipelineResultOutput {
	return o
}

func (o LookupImportPipelineResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupImportPipelineResult] {
	return pulumix.Output[LookupImportPipelineResult]{
		OutputState: o.OutputState,
	}
}

// The resource ID.
func (o LookupImportPipelineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportPipelineResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the import pipeline.
func (o LookupImportPipelineResultOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupImportPipelineResult) *IdentityPropertiesResponse { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

// The location of the import pipeline.
func (o LookupImportPipelineResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImportPipelineResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o LookupImportPipelineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportPipelineResult) string { return v.Name }).(pulumi.StringOutput)
}

// The list of all options configured for the pipeline.
func (o LookupImportPipelineResultOutput) Options() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImportPipelineResult) []string { return v.Options }).(pulumi.StringArrayOutput)
}

// The provisioning state of the pipeline at the time the operation was called.
func (o LookupImportPipelineResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportPipelineResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The source properties of the import pipeline.
func (o LookupImportPipelineResultOutput) Source() ImportPipelineSourcePropertiesResponseOutput {
	return o.ApplyT(func(v LookupImportPipelineResult) ImportPipelineSourcePropertiesResponse { return v.Source }).(ImportPipelineSourcePropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupImportPipelineResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupImportPipelineResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The properties that describe the trigger of the import pipeline.
func (o LookupImportPipelineResultOutput) Trigger() PipelineTriggerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupImportPipelineResult) *PipelineTriggerPropertiesResponse { return v.Trigger }).(PipelineTriggerPropertiesResponsePtrOutput)
}

// The type of the resource.
func (o LookupImportPipelineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportPipelineResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupImportPipelineResultOutput{})
}
