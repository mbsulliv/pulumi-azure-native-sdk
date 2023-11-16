// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the properties of the export pipeline.
func LookupExportPipeline(ctx *pulumi.Context, args *LookupExportPipelineArgs, opts ...pulumi.InvokeOption) (*LookupExportPipelineResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupExportPipelineResult
	err := ctx.Invoke("azure-native:containerregistry/v20230101preview:getExportPipeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExportPipelineArgs struct {
	// The name of the export pipeline.
	ExportPipelineName string `pulumi:"exportPipelineName"`
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An object that represents an export pipeline for a container registry.
type LookupExportPipelineResult struct {
	// The resource ID.
	Id string `pulumi:"id"`
	// The identity of the export pipeline.
	Identity *IdentityPropertiesResponse `pulumi:"identity"`
	// The location of the export pipeline.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The list of all options configured for the pipeline.
	Options []string `pulumi:"options"`
	// The provisioning state of the pipeline at the time the operation was called.
	ProvisioningState string `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The target properties of the export pipeline.
	Target ExportPipelineTargetPropertiesResponse `pulumi:"target"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupExportPipelineOutput(ctx *pulumi.Context, args LookupExportPipelineOutputArgs, opts ...pulumi.InvokeOption) LookupExportPipelineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExportPipelineResult, error) {
			args := v.(LookupExportPipelineArgs)
			r, err := LookupExportPipeline(ctx, &args, opts...)
			var s LookupExportPipelineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExportPipelineResultOutput)
}

type LookupExportPipelineOutputArgs struct {
	// The name of the export pipeline.
	ExportPipelineName pulumi.StringInput `pulumi:"exportPipelineName"`
	// The name of the container registry.
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExportPipelineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExportPipelineArgs)(nil)).Elem()
}

// An object that represents an export pipeline for a container registry.
type LookupExportPipelineResultOutput struct{ *pulumi.OutputState }

func (LookupExportPipelineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExportPipelineResult)(nil)).Elem()
}

func (o LookupExportPipelineResultOutput) ToLookupExportPipelineResultOutput() LookupExportPipelineResultOutput {
	return o
}

func (o LookupExportPipelineResultOutput) ToLookupExportPipelineResultOutputWithContext(ctx context.Context) LookupExportPipelineResultOutput {
	return o
}

// The resource ID.
func (o LookupExportPipelineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportPipelineResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the export pipeline.
func (o LookupExportPipelineResultOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupExportPipelineResult) *IdentityPropertiesResponse { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

// The location of the export pipeline.
func (o LookupExportPipelineResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExportPipelineResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o LookupExportPipelineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportPipelineResult) string { return v.Name }).(pulumi.StringOutput)
}

// The list of all options configured for the pipeline.
func (o LookupExportPipelineResultOutput) Options() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupExportPipelineResult) []string { return v.Options }).(pulumi.StringArrayOutput)
}

// The provisioning state of the pipeline at the time the operation was called.
func (o LookupExportPipelineResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportPipelineResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupExportPipelineResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupExportPipelineResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The target properties of the export pipeline.
func (o LookupExportPipelineResultOutput) Target() ExportPipelineTargetPropertiesResponseOutput {
	return o.ApplyT(func(v LookupExportPipelineResult) ExportPipelineTargetPropertiesResponse { return v.Target }).(ExportPipelineTargetPropertiesResponseOutput)
}

// The type of the resource.
func (o LookupExportPipelineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportPipelineResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExportPipelineResultOutput{})
}
